import sys

iana_tlds = sys.argv[1]
removed_tlds = sys.argv[2]
out_file = sys.argv[3]

tlds = set()

for l in open(iana_tlds, 'r'):
    l = l.strip()
    if l[0] == '#':
        continue
    tlds.add(l)

first = True 
for l in open(removed_tlds, 'r'):
    if first == True:
        first = False 
        continue
    l = l.strip().split(',')
    tlds.add(l[0].upper())

with open(out_file, 'w') as ofile:
    ofile.write("package util\n")
    ofile.write("var tldMap = map[string]bool{\n")
    for tld in tlds:
        ofile.write('"'+tld+'": true,\n')
    ofile.write("}")
