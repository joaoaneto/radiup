import argparse

def gen_basic(entities):

	"""
	Get the entities in a dict and generate according to Go Lang syntax
	"""
	print("Generating basic entities for Go Lang...")
	for i in entities:
		print("type %s struct {\n    %s \n}" % (i[0],'\n    '.join([str(k + ' ' + v) for k, v in i[1].items()])))

def main():

	# make the parser for command-line options
	parser = argparse.ArgumentParser()
	parser.add_argument('-g', '--gorm', action='store_true', help='gorm entities flag')
	parser.add_argument('-b', '--basic', action='store_true', help='basic entities flag')
	args = parser.parse_args()

	User = ["User", {"name":"string", "username":"string", "birthDate":"string", "email":"string"}]
	Music = ["Music", {"name":"string", "artist":"[]string", "id":"string"}]

	entities = [User, Music]

	if args.basic:
		gen_basic(entities)

if __name__ == "__main__":
	main()