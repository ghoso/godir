#Godir

Simple program which show directory information by JSON format.

Usage: godir [directory]

Output example: 
```json
{
    "IsDir": true,
    "Path": "/home/user1/test",
    "Name": "test",
    "Children": {
        "test1d": {
            "IsDir": true,
            "Path": "/home/user1/test/test1d",
            "Name": "test1d",
            "Children": {
                "hoge.txt": {
                    "IsDir": false,
                    "Path": "/home/user1/test/test1d/hoge.txt",
                    "Name": "hoge.txt",
                    "Children": null
                }
            }
        },
        "foo.txt": {
            "IsDir": true,
            "Path": "/home/user1/test/test1d",
            "Name": "test1d",
            "Children": null
        }
    }
}
```
