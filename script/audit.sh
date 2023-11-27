#!/bin/bash
Continue() {
	echo "Press any key to continue..."
	while true; do
		read -t 3 -n 1
		if [ $? = 0 ]; then
			clear
			break
		fi
	done
}
clear
echo "Try to run ./push-swap."
./build/push-swap
Continue
echo "Try to run ./push-swap '2 1 3 6 5 8' "
./build/push-swap "2 1 3 6 5 8"
Continue
echo "Try to run ./push-swap '0 1 2 3 4 5' "
./build/push-swap "0 1 2 3 4 5"
Continue
echo "Try to run ./push-swap '0 one 2 3' "
./build/push-swap "0 one 2 3"
Continue
echo "Try to run ./push-swap '1 2 2 3' "
./build/push-swap "1 2 2 3"
Continue
echo "Try to run ./push-swap '<5 random numbers>' with 5 random numbers instead of the tag."
read -t 10 -p "> " input_numbers
# echo "Vous avez entré : $input_numbers"
./build/push-swap "$input_numbers"
Continue
echo "Try to run ./push-swap '<5 random numbers>' with 5 different random numbers instead of the tag."
read -t 10 -p "> " input_numbers
./build/push-swap "$input_numbers"
Continue
echo "Try to run ./Checker and input nothing."
./build/Checker
Continue
echo "Try to run ./Checker '0 one 2 3' "
./build/Checker "0 one 2 3"
Continue
echo "Try to run echo -e 'sa\npb\nrrr\n' | ./Checker '0 9 1 8 2 7 3 6 4 5'"
echo -e "sa\npb\nrrr\n" | ./build/Checker "0 9 1 8 2 7 3 6 4 5"
Continue
echo "Try to run echo -e 'pb\nra\npb\nra\nsa\nra\npa\npa\n' | ./Checker '0 9 1 8 2'"
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./build/Checker "0 9 1 8 2"
Continue
echo "Try to run ARG=('4 67 3 87 23'); ./push-swap ARG | ./Checker ARG"
ARG=('4 67 3 87 23')
./build/push-swap "$ARG" | ./build/Checker "$ARG"
Continue
echo "Try to run ARG=('<100 random numbers>'); ./push-swap ARG with 100 random different numbers instead of the tag."
Continue
read -t 10 -p "> " input_numbers
./build/push-swap "$input_numbers"
Continue
echo "Try to run ARG=('<100 random numbers>'); ./push-swap ARG | ./Checker ARG with 100 random different numbers instead of the tag."
read -t 10 -p "> " input_numbers
./build/push-swap "$input_numbers"
Continue
