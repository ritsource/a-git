# Test
go test src/commit_test.go > /dev/null 2>&1
if [ $? -eq 0 ]; then
    RESULT=true
else
    RESULT=false
fi


# Printing result
echo -e "\e[1mUnit tests \"commit_test.go\""
if [ "$RESULT" != false ]; then
    echo -e "\e[32m\e[1mPASS :)"
else
    echo -e "\e[31m\e[1mFAIL :("
fi

echo -e "\e[0m"

rm -rf .git