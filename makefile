SRC = solution/main.cpp
DESTINATION_BUILD = solution/main.exe

run:	
	g++ ${SRC} -o ${DESTINATION_BUILD}
	./${DESTINATION_BUILD}