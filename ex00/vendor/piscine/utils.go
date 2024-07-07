package piscine

// remove fmt
import "ft"

const longLongUsage = "\033[1m\033[37m################## \033[5m\033[97mUsage\033[m\033[1m\033[37m ######################\n#                                             #\n#  Takes 1 argument, the name of a file       #\n#     satisfies the following conditions.     #\n#                                             #\n#  1. This file contains 1-26 Tetriminos.     #\n#                                             #\n#  2. Every Tetrimino consists of             #\n#      4 block chars, which is located in     #\n#      \033[90m4x4 square\033[37m(4 lines by 4 chars),        #\n#      each followed by a new line.           #\n#                                             #\n#  3. Each char is either a \033[90mblock char\033[37m (\033[5m'#'\033[m\033[1m\033[37m)  #\n#       or an \033[90mempty char\033[37m (\033[5m'.'\033[m\033[1m\033[37m).               #\n#                                             #\n#  4. Each block of a Tetrimino touchs        #\n#      at least 1 other block                 #\n#      on any of its 4 sides                  #\n#      (up, down, left and right).            #\n#                                             #\n#  5. A rotated Tetrimino is a different      #\n#      Tetrimino from the original.           #\n#                                             #\n###############################################\n\033[m"

// 	write there something like print string
// all start with ft
func Ft_PutStr(s string) {
	for _, c := range(s) {
		ft.PrintRune(c)
	}
}

func Ft_PutStrLn(s string) {
	for _, c := range(s) {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func Ft_PutUsage() {
	Ft_PutStr(longLongUsage)
}

func Ft_PrintErr() {
	Ft_PutStrLn("error")
}
