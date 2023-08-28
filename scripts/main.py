from rembg import remove

input_path = 'rock.jpeg'
output_path = 'output.png'

def erase(path):
    with open(path, 'rb') as i:
        with open(output_path, 'wb') as o:
            input = i.read()
            output = remove(input)
            o.write(output)
            return output_path
        