# Initial repository
git init > /dev/null 2>&1

# Creating a new file
echo 'Hello World' >foo.txt

# Adding files into tracking
git add . > /dev/null 2>&1

# aking a commit
git commit -m "Hello World" > /dev/null 2>&1

# reading blob content with catfile command
# 557db03de997c86a4a028e1ebd3a1ceb225be238 is the cnreypted sha for "Hello World"
OUTPUT="$(./a-git cat-file -p 557db03de997c86a4a028e1ebd3a1ceb225be238)"

# Checking OUTPUT
if [ "${OUTPUT}" == "Hello World" ]; then
    RESULT=true
else
    RESULT=false
fi

# Printing RESULT
echo -e "\e[1mTesting \"a-git cat-file\" cmd"
if [ "$RESULT" != false ]; then
    echo -e "\e[32m\e[1mPASS :)"
else
    echo -e "\e[31m\e[1mFAIL :("
fi

# Empty line
echo -e "\e[0m"

# Deleting repository 
rm -rf git