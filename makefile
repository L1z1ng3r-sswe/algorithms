SRC = main.cpp
DESTINATION_BUILD = main.exe

run:
	g++ ${SRC} -o ${DESTINATION_BUILD}
	./main.exe