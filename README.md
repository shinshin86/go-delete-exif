# go-delete-exif
Command line tool to remove Exif information from JPEG images.  
This is code I created to learn Go and may not be practical.

## Install

```
go install github.com/shinshin86/go-delete-exif
```

## Usage

```
go-delete-exif <input image path(JPEG only)>
```

If the process is successful, a file with `_exif-deleted.JPG` at the end is output.

### -d option

Add this option if you want to delete the target file after exif deletion.  
In that case, the file with `_exif-deleted.JPG` will not be generated.

```
go-delete-exif -d <input image path(JPEG only)>
```

## Description of process
Once converted to a PNG image, the exif information is removed.

```
JPEG ->      PNG      -> JPEG
        (Delete exif)
```