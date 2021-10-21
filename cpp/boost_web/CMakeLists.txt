cmake_minimum_required(VERSION 3.10) 

# set the project name 
project(web) 

set(CMAKE_CXX_STANDARD 17) 
set(CMAKE_CXX_STANDARD_REQUIRED True) 

set(BOOST_ROOT /usr/local/lib)
set(Boost_NO_SYSTEM_PATHS ON)

find_package(Boost 1.71.0 REQUIRED COMPONENTS system filesystem thread coroutine timer)

# add the executable 
if(Boost_FOUND)
	include_directories(${Boost_INCLUDE_DIRS})
	MESSAGE( STATUS "Boost_INCLUDE_DIRS = ${Boost_INCLUDE_DIRS}.")
	MESSAGE( STATUS "Boost_LIBRARIES = ${Boost_LIBRARIES}.")
	MESSAGE( STATUS "Boost_LIB_VERSION = ${Boost_LIB_VERSION}.")
	add_executable(web main.cc web.cc) 
	target_link_libraries (web ${Boost_LIBRARIES})
endif()
