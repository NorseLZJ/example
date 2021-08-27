cmake_minimum_required(VERSION 3.16)

project(norse)

set(CMAKE_CXX_STANDARD 17)

SET(EXECUTABLE_OUTPUT_PATH ${PROJECT_SOURCE_DIR}/build)

add_executable(norse main.cc env.cc)