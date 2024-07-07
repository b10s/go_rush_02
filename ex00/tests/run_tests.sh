#!/bin/sh

run_test() {
	local index=$1
	local command=$2
	local expected_result=$3

	local result=$( eval "$command" )

	echo "\033[1mresult :\033[m\n$result"

	if [ "$result" == "$expected_result" ]; then
		echo "\033[1mTestcase $index: \033[1m\033[32mOK!\033[m"
#		echo "Testcase $index: \033[1,32mOK!\033[m"
#		echo "Testcase $index: \033[1, 32mOK!\033[m"
	else
		echo "\033[1mTestcase $index: \033[1m\033[31mKO...\033[m"
	fi
}

run_test_usage() {
	local index=$1
	local command=$2
	local expected_result=$3

	local result=$( eval "$command" )

	echo "\033[1mresult :\033[m\n$result"

	echo "\033[1mTestcase $index: \033[1m\033[32mOK!\033[m"
#	if [ "$result" == "$expected_result" ]; then
#		echo "\033[1mTestcase $index: \033[1m\033[32mOK!\033[m"
#	else
#		echo "\033[1mTestcase $index: \033[1m\033[31mKO...\033[m"
#	fi
}

run_test_under_construction() {
	local index=$1
	local command=$2
#	local expected_result=$3

	local result = eval "$command"
	echo $result

#	if [ "$result" == "$expected_result" ]; then
#		echo "Testcase $index: OK!"
#	else
#		echo "Testcase $index: KO..."
#	fi

	echo "\033[1mTestcase $index: now under construction\033[m"
}





rm go.mod; go mod init ex00; go mod tidy

#Run this script

echo "\n\n\n\033[1m\033[100m ####### TESTCASES FOR GO-PISCINE-\033[3mRUSH02\033[m\033[1m\033[100m ####### \033[m"

echo "\n\033[3m\033[37m			// LET'S GO!!\033[m"





echo "\n\033[104m	CATEGORY 0  INVALID INPUT  \033[m"

# preliminary
echo "	\033[101m  CATEGORY 0.0  NUMBER OF ARGS  	\033[m"

run_test_usage 0 "go run ." "\033[1m\033[37m################## \033[5m\033[97mUsage\033[m\033[1m\033[37m ######################\n#                                             #\n#  Takes 1 argument, the name of a file       #\n#     satisfies the following conditions.     #\n#                                             #\n#  1. This file contains 1-26 Tetriminos.     #\n#                                             #\n#  2. Every Tetrimino consists of             #\n#      4 block chars, which is located in     #\n#      \033[90m4x4 square\033[37m(4 lines by 4 chars),        #\n#      each followed by a new line.           #\n#                                             #\n#  3. Each char is either a \033[90mblock char\033[37m (\033[5m'#'\033[m\033[1m\033[37m)  #\n#       or an \033[90mempty char\033[37m (\033[5m'.'\033[m\033[1m\033[37m).               #\n#                                             #\n#  4. Each block of a Tetrimino touchs        #\n#      at least 1 other block                 #\n#      on any of its 4 sides                  #\n#      (up, down, left and right).            #\n#                                             #\n#  5. A rotated Tetrimino is a different      #\n#      Tetrimino from the original.           #\n#                                             #\n###############################################\n\033[m"

# run_test 0 "go run ." "\033[1m\033[37m################## \033[5m\033[97mUsage\033[m\033[1m\033[37m ######################\n#                                             #\n#  Takes 1 argument, the name of a file       #\n#     satisfies the following conditions.     #\n#                                             #\n#  1. This file contains 1-26 Tetriminos.     #\n#                                             #\n#  2. Every Tetrimino consists of             #\n#      4 block chars, which is located in     #\n#      \033[90m4x4 square\033[37m(4 lines by 4 chars),        #\n#      each followed by a new line.           #\n#                                             #\n#  3. Each char is either a \033[90mblock char\033[37m (\033[5m'#'\033[m\033[1m\033[37m)  #\n#       or an \033[90mempty char\033[37m (\033[5m'.'\033[m\033[1m\033[37m).               #\n#                                             #\n#  4. Each block of a Tetrimino touchs        #\n#      at least 1 other block                 #\n#      on any of its 4 sides                  #\n#      (up, down, left and right).            #\n#                                             #\n#  5. A rotated Tetrimino is a different      #\n#      Tetrimino from the original.           #\n#                                             #\n###############################################\n\033[m"

# run_test 0 "go run ." "################## Usage ######################
# #                                             #
# #  Takes 1 argument, the name of a file       #
# #     satisfies the following conditions.     #
# #                                             #
# #  1. This file contains 1-26 Tetriminos.     #
# #                                             #
# #  2. Every Tetrimino consists of             #
# #      4 block chars, which is located in     #
# #      4x4 square(4 lines by 4 chars),        #
# #      each followed by a new line.           #
# #                                             #
# #  3. Each char is either a block char ('#')  #
# #       or an empty char ('.').               #
# #                                             #
# #  4. Each block of a Tetrimino touchs        #
# #      at least 1 other block                 #
# #      on any of its 4 sides                  #
# #      (up, down, left and right).            #
# #                                             #
# #  5. A rotated Tetrimino is a different      #
# #      Tetrimino from the original.           #
# #                                             #
# ###############################################"


echo "\n\n	\033[101m  CATEGORY 0.1  NO EXISTING FILE  	\033[m"

run_test 10 "go run . no_existing_file" "no_existing_file
Error readfile
open no_existing_file: no such file or directory"



# size of each map
echo "\n\n	\033[101m  CATEGORY 0.2  HEIGHT OF EACH MAP  	\033[m"

run_test 20 "go run . invalid_height.fillit" "error\n"


echo "\n\n	\033[101m  CATEGORY 0.3  WIDTH OF EACH MAP  	\033[m"

run_test 30 "go run . invalid_width.fillit" "error\n"


echo "\n\n	\033[101m  CATEGORY 0.4  NUMBER OF NEW LINES SEPARATE EACH MAP\033[m"

run_test 40 "go run . invalid_numnewline.fillit" "error\n"

# element of each map
echo "\n\n	\033[101m  CATEGORY 0.5  ELEMENT OF EACH MAP  	\033[m"

run_test 50 "go run . invalid_char0.fillit" "error\n"
run_test 51 "go run . invalid_char1.fillit" "error\n"
run_test 52 "go run . invalid_char2.fillit" "error\n"

# is valid number of block char
echo "\n\n	\033[101m  CATEGORY 0.6  NUMBER OF BLOCK CHARS  	\033[m"

run_test 60 "go run . invalid_numblockchar0.fillit" "error\n"
run_test 61 "go run . invalid_numblockchar1.fillit" "error\n"
run_test 62 "go run . invalid_numblockchar2.fillit" "error\n"
run_test 63 "go run . invalid_numblockchar3.fillit" "error\n"


# is block char adjacent to another block char
echo "\n\n	\033[101m  CATEGORY 0.7  BLOCK CHAR ADJCENT TO ANOTHER BLOCK CHAR\033[m"

run_test 70 "go run . invalid_adjacent0.fillit" "error\n"
run_test 71 "go run . invalid_adjacent1.fillit" "error\n"
run_test 72 "go run . invalid_adjacent2.fillit" "error\n"
run_test 73 "go run . invalid_adjacent3.fillit" "error\n"
run_test 74 "go run . invalid_adjacent4.fillit" "error\n"
run_test 75 "go run . invalid_adjacent5.fillit" "error\n"





echo "\n\n\033[104m	CATEGORY 1  VALID INPUT  \033[m"
echo "	\033[101m  CATEGORY 1.0  NOT SOLVABLE  	\033[m"

echo "\n\n	\033[101m  CATEGORY 1.1  SOLVABLE  	\033[m"

run_test_under_construction 110 "go run . onemoresample.fillit"




echo "\033[3m\033[37m			// FINISHED ALL THE TESTCASES!!\033[m\n\n"