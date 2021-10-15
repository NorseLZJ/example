cmake_minimum_required(VERSION 3.10) 

# set the project name 
project(web) 

set(CMAKE_CXX_STANDARD 17) 
set(CMAKE_CXX_STANDARD_REQUIRED True) 

link_libraries(pthread)

# add the executable 
add_executable(web main.cc) 

