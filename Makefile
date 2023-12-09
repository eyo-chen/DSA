# Makefile for your C++ program

# Compiler and compiler flags
CXX = g++
CXXFLAGS = -std=c++11 -Wall

# Source file and output file (default values)
SOURCE = main.cpp
OUTPUT = my_program

run:
	@if [ -z "$(file)" ]; then \
		echo "Usage: make run file=<source_file.cpp>"; \
	else \
		$(CXX) $(CXXFLAGS) -o $(OUTPUT) $(file); \
		./$(OUTPUT); \
		rm -f ./$(OUTPUT); \
	fi
