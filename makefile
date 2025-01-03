SRC = solution/main.cpp
DESTINATION_BUILD = solution/main

run:	
	g++ ${SRC} -o ${DESTINATION_BUILD}
	./${DESTINATION_BUILD}

	# go run solution/main.go