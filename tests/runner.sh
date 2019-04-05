# Prining strt
echo -e "\n\e[1mTesting..\e[0m"

# Line Break
echo -e "\e[0m"

# Run "a-git init"
bash tests/init.sh

# Run "a-git ls-tree"
# bash tests/lsTree.sh

# Run "a-git cat-file"
bash tests/catFile.sh

# Running unit tests inside docker
# tree_test.go
bash tests/unit_tests/tree_test.sh