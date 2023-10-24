#!/bin/bash

# Check if the script is being run as root
if [ "$(id -u)" != "0" ]; then
    echo "This script must be run as root."
    exit 1
fi

# Define the username and password
username="avmadmin"
password=$(openssl rand -base64 24)
echo $password

# Create the user
useradd -m "$username"

# Set the password (Note: This is not secure for production use!)
echo "$username:$password" | chpasswd

usermod -aG sudo $username

sudo -u $username ssh-keygen -t rsa -b 2048 -f /home/$username/.ssh/id_rsa -N ""

# Step 4: Copy the public key to the authorized_keys file
cat /home/$username/.ssh/id_rsa.pub >> /home/$username/.ssh/authorized_keys

echo "$username ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$username-nopasswd
systemctl restart ssh

echo ""
echo "Provide the below SSH key to your Avertium POC in a secure way"
echo ""
cat /home/$username/.ssh/id_rsa
