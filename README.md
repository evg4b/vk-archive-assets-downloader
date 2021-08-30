# vk-archive-assets-downloader

## Quick start:

1. Request and download archive on [this page](https://vk.com/data_protection?section=rules)
2. Unarchive to a folder
3. Run `./vk-archive-assets-downloader --src='~/downloads/archive'`

## Parameters
``` bash
./vk-archive-assets-downloader --src=~/downloads/archive --dialogs=259460211,359460212 --types=Photo,Video -dest ~/dest
# OR
./vk-archive-assets-downloader -src ~/downloads/archive -dialogs 259460211,359460212 -types Photo,Video -dest ~/dest
```

| Parameter | Default value | Description |
|---|---|---|
| --src | `./archive` | Path to archive folder |
| --dialogs | `N/A` | Comma separated string with dialog ids. Dialog id you can found in dialog URL (sel query parameter) `https://vk.com/im?sel=<Dialog Id>` |
| -types | `N/A` |  Comma separated string with the attachment type name. A type must be specified in the language of the archive. Example `-types Photo,Video` for english and `-types Фотография,Видеозапись` for russian |
| -dest | `./dest` | Path to destination folder. |