# Setting up GitHub with SSH authentication

## Check for existing SSH keys

You should check if you have any existing SSH keys before generating new ones.
You can do this by running the following command:

```sh
$ ls -la ~/.ssh
```

This lists all the content in the ´.ssh´ directory. The output may look like this:

```
drwx------  2 jostein jostein 4096 Aug  7  2023 .
drwxr-xr-x 62 jostein jostein 4096 Aug 23 12:09 ..
-rw-------  1 jostein jostein 2610 Jan 19  2024 id_rsa
-rw-r--r--  1 jostein jostein  577 Nov 14  2022 id_rsa.pub
-rw-r--r--  1 jostein jostein 5328 Jan  4  2024 known_hosts
```

If you see output like this you may skip to the part about [reading the contents of your generated key below](#reading-the-contents-of-your-generated-key).
## Generating a SSH keypair

To work with a repository using SSH you need to add a `public key` to your GitHub account. You can create a key pair (private and public key) using `ssh-keygen` in a terminal:

```sh
$ ssh-keygen
```

Running this command will give you output similar to this:

```sh
Generating public/private rsa key pair.
Enter file in which to save the key (/home/user/.ssh/id_rsa): # Press enter
Enter passphrase (empty for no passphrase): # No need for a password
Enter same passphrase again: # No need for a password
Your identification has been saved in id_rsa
Your public key has been saved in id_rsa.pub
The key fingerprint is:
< some string >
The key's randomart image is:
+---[RSA 3072]----+
|    oo*.+oo.     |
|   . *.=o=.      |
|  Eo+ .+o .      |
|=+o.. o o. . .   |
|*=+o . oS.+ +    |
|o*o   o  o + .   |
|o +  .    o .    |
| +         +     |
|  .       o..    |
+----[SHA256]-----+
```

When asked for the file in which to save the key, and when prompted for a password, you can just **press enter**.

### Reading the contents of your generated key

After having generated a key pair, the files, `id_rsa` and `id_rsa.pub` will be stored in a folder named `.ssh`. You need to read the `id_rsa.pub` file and copy its contents. In most cases reading the file can be done by running this command:

```sh
$ cat ~/.ssh/id_rsa.pub
```

The output from running this command is the contents of the file, and will look similar to this:

```sh
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDRfcVGuPHpt8iaZNRs9NeFGa+R4QbquwOKnIFoGJtqHbrzvWVfwcjfAsEV2/lu5yhMTmrp6lDtxXhTEqdoS/vLdQxtpVYqkDJd//HMyBeqk+oQTBHFy0K4axg1q+CMkigLe/woIb2AQn9xmjVS6RtqKyTVtMiO/sRXLqXBKtVmZXnKLHAF97JLxpzxWQznN/SgB3HCLQnxkPZbWEatAxXtc6zyO8mgA3iHpZ3OCJoNw5Z1ZMzmn0WXcawgcvrPrEYywPfGqXSi+uliOvnv/odwY85F3egJDNNuTuSCNyoxPN4E6C+2FvpQJD5mHzKuRh+8S3XD5FpVtpQ18KGcQI6L8H5NE2VPHaxodKuU1GZA00KgGfLl8qu3LD9eJW3taTm7S4xYDJKtNUK5rWivZL7q+IAXJHVdnvuDK9RJmP4RGSf3Bi0E76rQqGev1ZasB1DId26NWlg3QlCnlBpus0lMhNMj1Zg+1J2JE1me2W3AaYtFTy29bvDF5Yl+sEHFAtU= user@machinename
```

This string has to be copied and added to your GitHub account in the next step.

## Adding the key to your GitHub account

Log in to your GitHub account and navigate to your profile settings, then select `SSH and GPG keys` in the sidebar. Then click the green `New SSH key` button. Clicking [here](https://github.com/settings/ssh/new) should take you there directly.

In the `Title` textbox on this page you can name your key. Name it whatever you feel like.

In the `Key` textbox you ***need to paste the contents you copied from the `id_rsa.pub` file***.

Finally, click the `Add SSH key` button.

Now you should be able to clone GitHub repositories using SSH.

## Cloning a repository

Now that you have set up a SSH key on your GitHub account you should be ready to clone your repository: 

> NOTE: You need to replace `username` with your GitHub username in the snippet below

```sh
$ git clone git@github.com:dat320-2024/username-labs.git assignments
```

This should clone your repository and save it in a folder named `assignments`.

If you get a message similar to this:

```sh
The authenticity of host 'github.com (140.82.121.4)' can't be established. ED25519 key fingerprint is SHA256:+AAAAAAAAA/BBBBBBBBB. This key is not known by any other names Are you sure you want to continue connecting (yes/no/[fingerprint])?
```

type in `yes`, then press enter.


Now you should be able to navigate into your repository by running

```sh
$ cd assignments
```

If you have `Visual Studio Code` you should be able to open your repository with this command

```sh
$ code .
```

On Windows you can open your repository in the file explorer with this command:

```sh
$ explorer.exe .
```


## Adding the assignments to your repository

Your repository will most likely be empty when you clone it, unless you have made changes to it by other means.

To add the assignments that are published in the `assignments` repository you have to add that repository as a *remote*.

```sh
$ git remote add course-assignments git@github.com:dat320-2024/assignments.git
```

After having done so, you can tell `git` to fetch all files in the `assignments` repository and add it to your own repository with this command:

```sh
$ git pull course-assignments main # Requests everything from `assignments`, adds it to your repository
```

You can run this command everytime a new assignment is published.
