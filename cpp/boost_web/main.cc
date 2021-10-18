#include "web.h"

int main(int argc, char *argv[])
{
	webserver::Web web;

	web.run(argc, argv);
}