cmake_minimum_required(VERSION 3.17)
project(rep2)

set(CMAKE_CXX_STANDARD 17)

SET(EXECUTABLE_OUTPUT_PATH ${PROJECT_SOURCE_DIR})

set_property(GLOBAL PROPERTY RULE_MESSAGES OFF)

add_executable(rep2 main.cpp json.hpp)