#ifndef __WEB_H_
#define __WEB_H_

#include <boost/beast/core.hpp>
#include <boost/beast/http.hpp>
#include <boost/beast/version.hpp>
#include <boost/asio/ip/tcp.hpp>
#include <boost/asio/spawn.hpp>
#include <boost/config.hpp>
#include <algorithm>
#include <cstdlib>
#include <iostream>
#include <memory>
#include <string>
#include <thread>
#include <vector>
#include <unordered_map>
#include <functional>

namespace webserver
{
	namespace beast = boost::beast;	  // from <boost/beast.hpp>
	namespace http = beast::http;	  // from <boost/beast/http.hpp>
	namespace net = boost::asio;	  // from <boost/asio.hpp>
	using tcp = boost::asio::ip::tcp; // from <boost/asio/ip/tcp.hpp>

	// This is the C++11 equivalent of a generic lambda.
	// The function object is used to send an HTTP message.

	struct send_lambda
	{
		beast::tcp_stream &stream_;
		bool &close_;
		beast::error_code &ec_;
		net::yield_context yield_;

		send_lambda(
			beast::tcp_stream &stream,
			bool &close,
			beast::error_code &ec,
			net::yield_context yield)
			: stream_(stream), close_(close), ec_(ec), yield_(yield)
		{
		}

		template <bool isRequest, class Body, class Fields>
		void operator()(http::message<isRequest, Body, Fields> &&msg) const
		{
			// Determine if we should close the connection after
			close_ = msg.need_eof();

			// We need the serializer here because the serializer requires
			// a non-const file_body, and the message oriented version of
			// http::write only works with const messages.
			http::serializer<isRequest, Body, Fields> sr{msg};
			http::async_write(stream_, sr, yield_[ec_]);
		}
	};

	class Web
	{
	public:
		void run(int argc, char **argv);

		static beast::string_view mime_type(beast::string_view path);

		static std::string path_cat(beast::string_view base, beast::string_view path);

		template <class Body, class Allocator, class Send>
		static void handle_request(beast::string_view doc_root, http::request<Body, http::basic_fields<Allocator>> &&req, Send &&send);

		/**
		 * @brief Report a failure
		*/
		static void fail(beast::error_code ec, char const *what);

		/**
	    	* @brief Handles an HTTP server connection
		*/
		static void do_session(beast::tcp_stream &stream, std::shared_ptr<std::string const> const &doc_root, net::yield_context yield);

		/**
	    	* @brief Handles an HTTP server connection
		*/
		static void do_listen(net::io_context &ioc, tcp::endpoint endpoint, std::shared_ptr<std::string const> const &doc_root, net::yield_context yield);

		~Web()
		{
		}

	private:
		// all method register
		typedef std::unordered_map<std::string, std::function<int(int, int)>> MethodMap;
		MethodMap method_map;
	};

}

#endif
