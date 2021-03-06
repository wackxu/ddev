<h1>Buildkite Test Agent Setup</h1>

We are using [Buildkite](https://buildkite.com/drud) for Windows and macOS testing. The build machines and buildkite-agent must be set up before use.

## Windows Test Agent Setup:

0. Create the user "testbot" on the machine. The password should be the password of testbot@drud.com.
1. Install [chocolatey](https://chocolatey.org/)
2. Install golang/mysql-cli/make/git/docker-ce/nssm with `choco install -y git mysql-cli golang make docker-for-windows nssm GoogleChrome zip jq composer` (If docker-toolbox, use that instead; you may have to download the release separately to get correct version.)
3. Enable gd and curl extensions in /c/tools/php72/php.ini
3. Install bats: `git clone git://github.com/bats-core/bats-core; cd bats-core; git checkout v1.1.0; ./install.sh`
3. If a laptop, set the "lid closing" setting in settings to do nothing.
4. Set the "Sleep after time" setting in settings to never.
5. Install the buildkite-agent. Use the latest release from [github.com/buildkite/agent](https://github.com/buildkite/agent/releases). It should go in /c/buildkite-agent, with the buildkite-agent.exe in /c/buildkite-agent/bin and the config in /c/buildkite-agent.
6. Update the buildkite-agent.cfg with the *token* and *tags*. Tags will probably be like `"os=windows,osvariant=windows10pro,dockertype=dockerforwindows"` or `"os=windows,osvariant=windows10pro,dockertype=toolbox"`.
7. Set up the agent to [run as a service](https://buildkite.com/docs/agent/v3/windows#running-as-a-service):
    - __on the "Log On" tab in the services widget it must be set up to log in as the primary user of the machine, so it inherits environment variables and home directory.__
8. Set up the machine to [automatically log in on boot](https://www.cnet.com/how-to/automatically-log-in-to-your-windows-10-pc/).  Run netplwiz, provide the password for the main user, uncheck the "require a password to log in".
9. On Docker Toolbox systems, add a link to "Docker Quickstart Terminal" in C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp (see [link](http://www.thewindowsclub.com/make-programs-run-on-startup-windows)).
10. On Docker-for-windows systems, launch Docker. It will offer to reconfigure Hyper-V and do a restart.
11. On Docker-for-windows, configured the C: and other drives as shared to docker.
12. On Docker Toolbox systems, make sure that nested virtualization is enabled however you need to enable it.
13. Edit /c/ProgramData/git/config to `autocrlf: false` and verify that `git config --list` shows only autocrlf: false. 
14. Run `winpty docker run -it -p 80 busybox ls` to trigger the Windows Defender warning, and "allow access".
15. Try running .buildkite/sanetestbot.sh to check your work.
16. Change the name of the machine to something in keeping with current style. Maybe `testbot-dell-toolbox-3`.
17. Reboot the machine and do a test run. (On windows the machine name only takes effect on reboot.)
18. Set the timezone properly (US MT)
19. Log into Chrome with the user testbot@drud.com and enable Chrome Remote Desktop.

### macOS Test Agent Setup

0. Create the user "testbot" on the machine. The password should be the password of testbot@drud.com.
1. Install [homebrew](https://brew.sh/)
2. Install golang/git/docker with `brew install golang git buildkite-agent mariadb jq p7zip bats-core composer`
3. Install docker with `brew cask install docker`
4. If the xcode command line tools are not yet installed, install them with `xcode select --install`
5. Edit the buildkite-agent.cfg in /usr/local/etc/buildkite-agent.cfg to add the agent token and the tags. Tags will probably be like `"os=macos,osvariant=highsierra,dockertype=dockerformac"` - Also edit with `build-path="~/tmp/buildkite-agent/builds"`
6. Install nosleep `brew cask install nosleep`
7. Enable nosleep using its shortcut in the Mac status bar.
8. In nosleep Preferences, enable "Never sleep on AC Adapter", "Never sleep on Battery", and "Start nosleep utility on system startup".
9. Set up Mac to [automatically log in on boot](https://support.apple.com/en-us/HT201476).
10. Try running .buildkite/sanetestbot.sh to check your work.
11. Change the name of the machine to something in keeping with current style. Maybe `testbot-mbp2017-macos-3`.
12. Log into Chrome with the user testbot@drud.com and enable Chrome Remote Desktop.
13. Set the timezone properly (US MT)
14. Reboot the machine and do a test run.
