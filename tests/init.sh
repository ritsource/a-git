# Initialize repo with "a-git init" command
./a-git init

# Assert if repo has been created or not
git status
if [ $? -eq 0 ]; then
    echo -e "\e[32mPASS :)"
else
    echo -e "\e[31mFAIL :("
fi

echo -e "\e[0m==================>"

rm -rf .git
