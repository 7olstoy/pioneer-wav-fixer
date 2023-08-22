# pioneer-wav-fixer

## About
Many files from bandcamp has the "unsupported file type error" issue with old Pioneer XDJ/CDJ decks.
This is because of WAV_EXT flag into these files. 

For solving this we need to rewrite 20th and 21th bytes of these wav-files to ```01 00``` and the Pioneer old decks will play the files.

## How to run
> :warning:
> **Make a backup of your files before running the tool!**

You can buld this project with this repo or use binary from [latest](https://github.com/7olstoy/pioneer-wav-fixer/releases/) release (amd64 and arm64 version).

#### MacOS

<details>
<summary>Download and allow the app</summary>

Open terminal and download the app:
```
curl -L https://github.com/7olstoy/pioneer-wav-fixer/releases/latest/download/pioneer_wav_fixer_$(uname -m) --output pioneer_wav_fixer
```
Add permission for execute:
```
chmod +x pioneer_wav_fixer
```
Run it:
```
./pioneer_wav_fixer help
```
Close the warning and open settings on your Mac, choose Apple menu > System Settings, then click Privacy & Security in the sidebar. After you need to approve the running on this app pioneer_wav_fixer and run it again:

<img title="warning" src="/images/macos-warning.png">

<img title="settings" src="/images/macos-settings.png">

```
./pioneer_wav_fixer help
```
If you see all options - you can continue.

</details>

Run the app with your files directory:
```
./pioneer_wav_fixer -folder /Volumes/USB_FLASH_NAME/Contents
```
and you will see all broken files. For fix it:
```
./pioneer_wav_fixer -folder /Volumes/USB_FLASH_NAME/Contents -overwrite
```

<details>
<summary>Options</summary>

```
  -folder string
    	Path to the folder containing WAV files (default ".")
  -list
    	Show all files and 20-21 bytes value
  -overwrite
    	Overwrite files with non-0100 20-21 bytes value, i.e. fix the main issue
```

</details>

## Thanks
Big up to pioneer [forum](https://forums.pioneerdj.com/hc/en-us/community/posts/360043048651-E-8305-unsupported-file-type-error) and [this](https://github.com/camm9909/WavPatcher) old python version.
