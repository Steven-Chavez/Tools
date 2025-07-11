"""
Ubuntu Linux SIEM Collector System Setup Script
============================

This script performs the following tasks:
1. Sets up automatic updates.
2. Configures Vim settings.
3. Creates rsyslog tuning configuration.
4. Creates a trash log for tuned logs.
5. Overwrites the logrotate configuration.
6. Opens port 514 for UDP/TCP.
7. Checks disk usage and ensures minimum requirements.
8. Restarts the rsyslog service.

Author: Steven Chavez
Date: 9/9/2024

Usage: python3 UbuntuLunix-SyslogCollectorSetup.py

"""
import subprocess

# Used for setting output color 
class Colors:
    RED = '\033[91m'
    GREEN = '\033[92m'
    YELLOW = '\033[93m'
    BLUE = '\033[94m'
    MAGENTA = '\033[95m'
    CYAN = '\033[96m'
    RESET = '\033[0m'

# Define the new content for the config file
logrotate_config = """
/var/log/syslog
/var/log/mail.info
/var/log/mail.warn
/var/log/mail.err
/var/log/mail.log
/var/log/daemon.log
/var/log/kern.log
/var/log/auth.log
/var/log/user.log
/var/log/lpr.log
/var/log/cron.log
/var/log/debug
/var/log/messages
/var/log/trash.log
{
        su root syslog
        rotate 1
        daily
        missingok
        notifempty
        compress
        create
        delaycompress
        sharedscripts
        postrotate
                /usr/lib/rsyslog/rsyslog-rotate
        endscript
}
"""


def create_trash_log():
    """
    Creates a trash log file for tuned logs.

    This function performs the following steps:
    1. Creates the /var/log/trash.log file.
    2. Changes the ownership of the file to syslog:adm.
    3. Sets the file permissions to 640.

    The function prints status messages to indicate the progress of each step.
    """
    
    print(f"{Colors.GREEN}#### CREATING TRASH LOG ####{Colors.RESET}")

    run_command("touch /var/log/trash.log", False)
    run_command("chown syslog:adm /var/log/trash.log", False)
    run_command("chmod 640 /var/log/trash.log", False)

    print("trash.log created to trash tuned logs...")
    print("")


def create_rsyslog_tuning_file():
    """
    Creates a tuning configuration file for rsyslog.

    This function performs the following steps:
    1. Creates the /etc/rsyslog.d/01-tuning.conf file.

    The function prints status messages to indicate the progress of each step.
    """
    
    print(f"{Colors.GREEN}#### CREATING RSYSLOG TUNNING CONFIG ####{Colors.RESET}")

    run_command("touch /etc/rsyslog.d/01-tuning.conf", False)

    print("01-tuning.conf was created to store rsyslog tunning commands...")
    print("")


def run_command(command, output):
    """
    Run a shell command and print its output.

    Parameters:
    command (str): The shell command to be executed.
    output (bool): If True, captures the output of the command.

    This function executes the given shell command using subprocess.run.
    If there is any error output (stderr), it prints the error message.
    """
    
    result = subprocess.run(command, shell=True, capture_output=output, text=False)
    if result.stderr:
        output = result.stderr
        if output != "None":
            print(output)


def setup_auto_update():
    """
    Sets up automatic updates for the system.

    This function performs the following steps:
    1. Updates the package lists.
    2. Upgrade all installed packages to their latest versions.
    3. Installs the 'unattended-upgrades' package to enable automatic updates.
    4. Checks the status of the 'unattended-upgrades' service to ensure it is active.

    The function uses the 'run_command' utility to execute system commands and prints
    status messages to indicate the progress of each step.
    """
    
    print(f"{Colors.GREEN}#### RUNNING UPDATES ####{Colors.RESET}")
    # Update package lists
    run_command("apt update", False)

    # Upgrade installed packages
    run_command("apt upgrade -y", False)

    print("")
    print(f"{Colors.GREEN}#### SETTING UP AUTOMATIC UPDATES ####{Colors.RESET}")
    # Install unattended-upgrades package
    run_command("apt install -y unattended-upgrades", False)

    print("")
    print(f"{Colors.GREEN}#### UNATTENDED UPGRADES STATUS ####{Colors.RESET}")
    # Check the status of the unattended-upgrades service
    run_command("systemctl status unattended-upgrades.service | grep 'Active'", False)

    print("")


def setup_vim_settings():
    """
    Configures Vim editor settings.

    This function sets up the following Vim configurations:
    1. Sets the color scheme to 'desert'.
    2. Enables the use of spaces instead of tabs.
    3. Sets the tab stop to 4 spaces.
    4. Sets the shift width to 4 spaces.
    5. Sets the soft tab stop to 4 spaces.
    6. Enables automatic indentation.

    The function uses the 'run_command' utility to append these settings to the '.vimrc' file.
    """
    
    run_command("echo 'color desert' > ~/.vimrc", False)
    run_command("echo 'set expandtab' >> ~/.vimrc", False)
    run_command("echo 'set tabstop=4' >> ~/.vimrc", False)
    run_command("echo 'set shiftwidth=4' >> ~/.vimrc", False)
    run_command("echo 'set softtabstop=4' >> ~/.vimrc", False)
    run_command("echo 'set autoindent' >> ~/.vimrc", False)


def overwrite_logrotate_config():
    """
    Overwrites the logrotate configuration for rsyslog.

    This function performs the following steps:
    1. Prints a status message indicating the start of the configuration process.
    2. Defines the path to the logrotate configuration file for rsyslog.
    3. Opens the specified file in write mode and overwrites it with new content.
    4. Prints a confirmation message once the configuration has been updated.

    The function uses the 'logrotate_config' variable to write the new configuration content.
    """
    
    print(f"{Colors.GREEN}#### CONFIGURING LOGROTATE CONFIG ####{Colors.RESET}")
    # Path to the config file
    file_path = "/etc/logrotate.d/rsyslog"

    # Open the file in write mode and overwrite it with the new content
    with open(file_path, 'w') as file:
        file.write(logrotate_config)

    print("/etc/logrotate.d/rsyslog config updated to rotate daily...")
    print("")


def open_port_514():
    """
    Opens port 514 for both UDP and TCP in the rsyslog configuration.

    This function performs the following steps:
    1. Prints a status message indicating the start of the process.
    2. Defines the path to the rsyslog configuration file.
    3. Reads the content of the rsyslog configuration file.
    4. Searches for and uncomments lines related to UDP and TCP port 514.
    5. Writes the updated content back to the configuration file.
    6. Prints a confirmation message once the port has been opened.

    The function modifies the rsyslog configuration to enable listening on port 514 for both UDP and TCP.
    """
    
    print(f"{Colors.GREEN}#### OPENING PORT 514 UDP/TCP ####{Colors.RESET}")
    # Path to the rsyslog.conf file
    file_path = "/etc/rsyslog.conf"

    # Read the file content
    with open(file_path, 'r') as file:
        lines = file.readlines()

    # Uncomment the lines related to UDP and TCP port 514
    with open(file_path, 'w') as file:
        for line in lines:
            if line.strip().startswith('#') and ('module(load="imudp")' in line or 'input(type="imudp" port="514")' in line or 'module(load="imtcp")' in line or 'input(type="imtcp" port="514")' in line):
                file.write(line.lstrip('#'))  # Remove the leading '#' to uncomment
            else:
                file.write(line)

    print("port 514 opened for udp and tcp...")
    print("")


def restart_rsyslog():
    """
    Restarts the rsyslog service.

    This function performs the following steps:
    1. Prints a status message indicating the start of the restart process.
    2. Executes the command to restart the rsyslog service using the 'run_command' utility.
    3. Prints a confirmation message once the rsyslog service has been restarted.

    The function ensures that the rsyslog service is restarted to apply any configuration changes.
    """
    
    print(f"{Colors.GREEN}#### RESTARTING RSYSLOG ####{Colors.RESET}")

    run_command("systemctl restart rsyslog", False)

    print("rsyslog restarted...")
    print("")


def get_disk_usage():
    """
    Checks the disk usage of the root filesystem.

    This function performs the following steps:
    1. Prints a status message indicating the start of the disk space check.
    2. Executes the 'df -h /' command to get the disk usage information.
    3. Parses the output to extract the disk size.
    4. Prints the disk size.
    5. Checks if the disk size is less than 250GB and prints a warning if it is.
    6. Prints a success message if the disk size meets the minimum requirement.

    The function ensures that the disk space meets the minimum requirement of 250GB.
    """
    
    print(f"{Colors.GREEN}#### CHECKING DISKSPACE ####{Colors.RESET}")

    # Run the df -h command
    result = subprocess.run(['df', '-h', '/'], stdout=subprocess.PIPE)
    output = result.stdout.decode('utf-8')

    # Parse the output
    lines = output.split('\n')
    values = lines[1].split()

    # Extract the relevant values
    size = values[1]
    print("Disk Size = " + size)

    size_int = float(size.replace("G", ""))

    if size_int < 250:
        print(f"{Colors.RED}WARNING{Colors.RESET} - Update diskspace to a minimum of 250GB")
    else:
        print(f"{Colors.CYAN}SUCCESS{Colors.RESET} - Disk space meets minimum requirements!")

    print("")


setup_auto_update()
setup_vim_settings()
create_rsyslog_tuning_file()
create_trash_log()
overwrite_logrotate_config()
open_port_514()
get_disk_usage()
restart_rsyslog()
