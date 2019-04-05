# Initialize repo with "a-git init" command
./a-git init > /dev/null 2>&1

# Assert if repo has been created or not
git status > /dev/null 2>&1
if [ $? -eq 0 ]; then
    RESULT=true
else
    RESULT=false
fi

# Printing result
echo -e "\e[1mTesting \"a-git init\" cmd"
if [ "$RESULT" == true ]; then
    echo -e "\e[32m\e[1mPASS :)"
else
    echo -e "\e[31m\e[1mFAIL :("
fi

# Echoing empty line
echo -e "\e[0m"

# Deletign created repo
rm -rf .git
