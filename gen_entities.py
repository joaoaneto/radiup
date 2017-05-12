import argparse

def gen_basic(entities):

	"""
	Get the entities in a dict and generate according to Go Lang syntax
	"""
	print("Generating basic entities for Go Lang...")
	for i in entities:
		print("type %s struct {\n    %s \n}" % (i[0],'\n    '.join([str(k + ' ' 
			+ v) for k, v in i[1].items()])))

def gen_mgo(entities):

	"""
	Get the entities in a dict and generate according to Go Lang syntax and mgo documentation
	http://gopkg.in/mgo.v2
	"""
	print("Generating mgo entities for Go Lang...")
	for i in entities:
		print("type %s struct {\n    %s \n}" % (i[0],'\n    '.join([str(k + ' ' + v + " `bson:\"" 
			+ k.lower() + "\"`") for k, v in i[1].items()])))	


def main():

	# make the parser for command-line options
	parser = argparse.ArgumentParser()
	parser.add_argument('-m', '--mgo', action='store_true', help='mgo entities flag')
	parser.add_argument('-b', '--basic', action='store_true', help='basic entities flag')
	args = parser.parse_args()

	User = ["User", {"name":"string", "username":"string", "birthDate":"string", "email":"string"}]
	Music = ["Music", {"name":"string", "artist":"[]string", "id":"string"}]

	entities = [User, Music]

	if args.basic:
		gen_basic(entities)
	elif args.mgo:
		gen_mgo(entities)
	else:
		print("Use -m for generate mgo entities or -b for generate basic entities")
		exit()

if __name__ == "__main__":
	main()