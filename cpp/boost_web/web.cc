#include "web.h"

namespace webserver
{

	void Web::run(int argc, char **argv)
	{
		// Check command line arguments.
		if (argc != 5)
		{
			std::cerr << "Usage: http-server-coro <address> <port> <doc_root> <threads>\n"
					  << "Example:\n"
					  << "    http-server-coro 0.0.0.0 8080 . 1\n";
			exit(EXIT_FAILURE);
		}
		auto const address = net::ip::make_address(argv[1]);
		auto const port = static_cast<unsigned short>(std::atoi(argv[2]));
		auto const doc_root = std::make_shared<std::string>(argv[3]);
		auto const threads = std::max<int>(1, std::atoi(argv[4]));

		// The io_context is required for all I/O
		net::io_context ioc{threads};

		// Spawn a listening port
		boost::asio::spawn(ioc,
						   std::bind(&do_listen, std::ref(ioc), tcp::endpoint{address, port}, doc_root, std::placeholders::_1));

		// Run the I/O service on the requested number of threads
		std::vector<std::thread> v;
		v.reserve(threads - 1);
		for (auto i = threads - 1; i > 0; --i)
			v.emplace_back([&ioc]
						   { ioc.run(); });
		ioc.run();

		exit(EXIT_SUCCESS);
	}

	beast::string_view Web::mime_type(beast::string_view path)
	{
		using beast::iequals;
		auto const ext = [&path]
		{
			auto const pos = path.rfind(".");
			if (pos == beast::string_view::npos)
				return beast::string_view{};
			return path.substr(pos);
		}();

		if (iequals(ext, ".htm"))
			return "text/html";
		if (iequals(ext, ".html"))
			return "text/html";
		if (iequals(ext, ".php"))
			return "text/html";
		if (iequals(ext, ".css"))
			return "text/css";
		if (iequals(ext, ".txt"))
			return "text/plain";
		if (iequals(ext, ".js"))
			return "application/javascript";
		if (iequals(ext, ".json"))
			return "application/json";
		if (iequals(ext, ".xml"))
			return "application/xml";
		if (iequals(ext, ".swf"))
			return "application/x-shockwave-flash";
		if (iequals(ext, ".flv"))
			return "video/x-flv";
		if (iequals(ext, ".png"))
			return "image/png";
		if (iequals(ext, ".jpe"))
			return "image/jpeg";
		if (iequals(ext, ".jpeg"))
			return "image/jpeg";
		if (iequals(ext, ".jpg"))
			return "image/jpeg";
		if (iequals(ext, ".gif"))
			return "image/gif";
		if (iequals(ext, ".bmp"))
			return "image/bmp";
		if (iequals(ext, ".ico"))
			return "image/vnd.microsoft.icon";
		if (iequals(ext, ".tiff"))
			return "image/tiff";
		if (iequals(ext, ".tif"))
			return "image/tiff";
		if (iequals(ext, ".svg"))
			return "image/svg+xml";
		if (iequals(ext, ".svgz"))
			return "image/svg+xml";
		return "application/text";
	}

	std::string Web::path_cat(beast::string_view base, beast::string_view path)
	{
		if (base.empty())
			return std::string(path);
		std::string result(base);
#ifdef BOOST_MSVC
		char constexpr path_separator = '\\';
		if (result.back() == path_separator)
			result.resize(result.size() - 1);
		result.append(path.data(), path.size());
		for (auto &c : result)
			if (c == '/')
				c = path_separator;
#else
		char constexpr path_separator = '/';
		if (result.back() == path_separator)
			result.resize(result.size() - 1);
		result.append(path.data(), path.size());
#endif
		return result;
	}

	template <class Body, class Allocator, class Send>
	void Web::handle_request(beast::string_view doc_root, http::request<Body, http::basic_fields<Allocator>> &&req, Send &&send)
	{
		// Returns a bad request response

		auto const bad_request = [&](beast::string_view why)
		{
			http::response<http::string_body> res{http::status::bad_request, req.version()};
			res.set(http::field::server, BOOST_BEAST_VERSION_STRING);
			res.set(http::field::content_type, "text/html");
			res.keep_alive(req.keep_alive());
			res.body() = std::string(why);
			res.prepare_payload();
			return res;
		};

		// Returns a not found response
		auto const not_found = [&](beast::string_view target)
		{
			http::response<http::string_body> res{http::status::not_found, req.version()};
			res.set(http::field::server, BOOST_BEAST_VERSION_STRING);
			res.set(http::field::content_type, "text/html");
			res.keep_alive(req.keep_alive());
			res.body() = "The resource '" + std::string(target) + "' was not found.";
			res.prepare_payload();
			return res;
		};

		// Returns a server error response
		auto const server_error = [&](beast::string_view what)
		{
			http::response<http::string_body> res{http::status::internal_server_error, req.version()};
			res.set(http::field::server, BOOST_BEAST_VERSION_STRING);
			res.set(http::field::content_type, "text/html");
			res.keep_alive(req.keep_alive());
			res.body() = "An error occurred: '" + std::string(what) + "'";
			res.prepare_payload();
			return res;
		};

		// Make sure we can handle the method
		switch (req.method())
		{
		case http::verb::get:
			break;
		default:
			return send(bad_request("Unknown HTTP-method"));
		}

		// Request path must be absolute and not contain "..".
		if (req.target().empty() || req.target()[0] != '/' || req.target().find("..") != beast::string_view::npos)
			return send(bad_request("Illegal request-target"));

		// Build the path to the requested file
		std::string path = path_cat(doc_root, req.target());
		std::cout << "REQ TARGRT ---> " << req.target() << std::endl;
		if (req.target().back() == '/')
			path.append("index.html");

		//MethodMap::iterator it = method_map.find(req.target());
		//if (it == this->method_map.end())
		//{
		//	return send(bad_request("Illegal request-target"));
		//}

		//it->second(req);

		// Attempt to open the file
		beast::error_code ec;
		http::file_body::value_type body;
		body.open(path.c_str(), beast::file_mode::scan, ec);

		// Handle the case where the file doesn't exist
		if (ec == beast::errc::no_such_file_or_directory)
			return send(not_found(req.target()));

		// Handle an unknown error
		if (ec)
			return send(server_error(ec.message()));

		// Cache the size since we need it after the move
		auto const size = body.size();

		// Respond to HEAD request
		if (req.method() == http::verb::head)
		{
			http::response<http::empty_body> res{http::status::ok, req.version()};
			res.set(http::field::server, BOOST_BEAST_VERSION_STRING);
			res.set(http::field::content_type, mime_type(path));
			res.content_length(size);
			res.keep_alive(req.keep_alive());
			return send(std::move(res));
		}

		// Respond to GET request
		http::response<http::file_body> res{
			std::piecewise_construct,
			std::make_tuple(std::move(body)),
			std::make_tuple(http::status::ok, req.version())};
		res.set(http::field::server, BOOST_BEAST_VERSION_STRING);
		res.set(http::field::content_type, mime_type(path));
		res.content_length(size);
		res.keep_alive(req.keep_alive());
		return send(std::move(res));
	}

	void Web::fail(beast::error_code ec, char const *what)
	{
		std::cerr << what << ": " << ec.message() << "\n";
	}

	void Web::do_session(beast::tcp_stream &stream, std::shared_ptr<std::string const> const &doc_root, net::yield_context yield)
	{
		bool close = false;
		beast::error_code ec;

		// This buffer is required to persist across reads
		beast::flat_buffer buffer;

		// This lambda is used to send messages
		send_lambda lambda{stream, close, ec, yield};

		for (;;)
		{
			// Set the timeout.
			stream.expires_after(std::chrono::seconds(30));

			// Read a request
			http::request<http::string_body> req;
			http::async_read(stream, buffer, req, yield[ec]);
			if (ec == http::error::end_of_stream)
				break;
			if (ec)
				return fail(ec, "read");

			// Send the response
			handle_request(*doc_root, std::move(req), lambda);
			if (ec)
				return fail(ec, "write");
			if (close)
			{
				// This means we should close the connection, usually because
				// the response indicated the "Connection: close" semantic.
				break;
			}
		}

		// Send a TCP shutdown
		stream.socket().shutdown(tcp::socket::shutdown_send, ec);

		// At this point the connection is closed gracefully
	}

	void Web::do_listen(net::io_context &ioc, tcp::endpoint endpoint, std::shared_ptr<std::string const> const &doc_root, net::yield_context yield)
	{
		beast::error_code ec;

		// Open the acceptor
		tcp::acceptor acceptor(ioc);
		acceptor.open(endpoint.protocol(), ec);
		if (ec)
			return fail(ec, "open");

		// Allow address reuse
		acceptor.set_option(net::socket_base::reuse_address(true), ec);
		if (ec)
			return fail(ec, "set_option");

		// Bind to the server address
		acceptor.bind(endpoint, ec);
		if (ec)
			return fail(ec, "bind");

		// Start listening for connections
		acceptor.listen(net::socket_base::max_listen_connections, ec);
		if (ec)
			return fail(ec, "listen");

		for (;;)
		{
			tcp::socket socket(ioc);
			acceptor.async_accept(socket, yield[ec]);
			if (ec)
			{
				fail(ec, "accept");
			}
			else
			{
				boost::asio::spawn(
					acceptor.get_executor(),
					std::bind(do_session, beast::tcp_stream(std::move(socket)), doc_root, std::placeholders::_1));
			}
		}
	}
}
