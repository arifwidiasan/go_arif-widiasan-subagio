#!/bin/sh
folder1="$1 at $(date)"
mkdir "$folder1"
cd "$folder1"
mkdir about_me
cd about_me
mkdir personal
cd personal
file1="facebook.txt"
echo "https://www.facebook.com/$2" > $file1
cd ..
mkdir professional
cd professional
file2="linkedin.txt"
echo "https://www.linkedin.com/in/$3" > $file2
cd ..
cd ..
mkdir my_friends
cd my_friends
file3="list_of_my_friend.txt"
curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > $file3
cd ..
mkdir my_system_info
cd my_system_info
file4="about_this_laptop.txt"
echo "My username: $USER" > $file4
host=$(uname -a)
echo "With host: $host" >> $file4
file5="internet_connection.txt"
echo "connection to google.com:" > $file5
echo "$(sudo ping google.com -c 3)" >> $file5

