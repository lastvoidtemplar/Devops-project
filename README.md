# Devops-project

## App

The app is a simple go app that is using the go standard library. It has two endpoints:

- / - returns "Hello World"
- /hostname - returns the hostname of the machine(it uses the environment HOSTNAME).

## Github Actions

There is a Github Actions workflow that builds the Docker Image on a  pull request or a push to main.

## OpenTofu

### Setup OpenTofu

Install tofuenv:

```bash
sudo apt update -y
sudo apt install -y jq gnupg
# Install tofuenv to local path $HOME/.tofuenv
git clone --depth=1 https://github.com/tofuutils/tofuenv.git ~/.tofuenv

# Add tofuenv path(bash)
echo 'export PATH="$HOME/.tofuenv/bin:$PATH"' >> ~/.bash_profile
# ---or---
# Add tofuenv path(zsh)
echo 'export PATH="$HOME/.tofuenv/bin:$PATH"' >> ~/.zshrc
```

Install tofu, choose a version and check the installation:

```bash
tofuenv install latest
tofuenv use latest
```

### Run OpenTofu

Setup your AWS credentials:

```bash
export AWS_ACCESS_KEY_ID="your AWS access key"
export AWS_SECRET_ACCESS_KEY="your AWS secret key"
```

Run OpenTofu:

```bash
cd terraform
tofu init
tofu plan
tofu apply
cd ..
```

To destroy the things created by OpenTofu, run:

```bash
cd terraform
tofu destroy
cd ..
```

## Ansible

### Setup Ansible

```bash
cd ansible
python3 -m venv venv
source venv/bin/activate
pip3 install -r requirements.txt
```

### Run Ansible

```bash
ansible-playbook deploy.yml -i inventory/ --key-file ~/.ssh/id_rsa_tofu
```
