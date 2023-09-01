import os
import argparse
from dotenv import load_dotenv
from rembg import remove

load_dotenv()
parser = argparse.ArgumentParser(
    prog="bgeraser-script",
    description="Script responsible for removing images backgrounds",
)
parser.add_argument("-i", "--input")
parser.add_argument("-o", "--output")


def erase(input_path, output_path):
    with open(input_path, "rb") as i:
        with open(output_path, "wb") as o:
            input = i.read()
            output = remove(input)
            o.write(output)
            return output_path


def main():
    print("[rembg] Starting script")
    try:
        args = parser.parse_args()

        if args.input is None or args.output is None:
            print("Both --input and --output are required arguments")
            os._exit(1)

        local_path = os.getenv("STORAGE_LOCAL_PATH")
        if local_path is None:
            print("Missing STORAGE_LOCAL_PATH variable")
            os._exit(1)

        input_path = local_path + "/" + args.input
        output_path = local_path + "/" + args.output
        erase(input_path, output_path)

    except Exception as err:
        print("Failed to remove image's background: ", err)
        os._exit(1)


main()
