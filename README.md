# Advent of Code 2020 solutions

Go-spaget.

```
make
./run.sh
```

## Solution descriptions

1. Brute-force search.
2. Brute-force search, count given character and compare to rules.
3. Trace line from beginning to bottom using given x and y offsets and count seen obstacle characters. Wrap around horizontal axis with modulo.
4. Tedious regular expression spaghetti.
5. Parse each line as binary numbers where each character `B` or `R` is set to 1, else to 0.
Find minimum and maximum numbers, then find the only ID missing between the min-max range.
6. Use 26 bits as a set for each ASCII lowercase character `a-z`, where bit 1 means character is present and 0 means absent. Compute union with bitwise or, intersection with bitwise and, then size of the sets with popcount.
7. Construct a directed graph where each node is a bag and each edge weight denotes how many bags that bag contains. Part A, search amount of unique nodes reachable from node `shiny-gold`. Part B, start from node `shiny-gold` and sum amount of child nodes, multiplied by edge weights.
8. Brute-force search, try running the program in an interpreter and do an early exit if there is a loop.
9. Brute-force search.
10. Compute all valid combination counts with dynamic programming.
11. Two 2D grids consisting of `byte` elements. Copy the grid on each update and continue updating until there are no changes.
12. Compute integer rotations with sin and cos in 2D space while carefully avoiding floating-point rounding errors.
13. Chinese remainder theorem and Go standard library big integers.
14. Combination of string parsing and bit twiddling according to the given rules. Use hash tables to simulate the memory to allow large indexes.
15. Brute-force search over all age steps `t` starting at 0, with hash tables of ages of each number.
16. Filter all valid tickets and sum invalid values (part A). Find all valid rules for every column/ticket-index. Select rules for columns by greedy search by assigning a rule to a column which has only one valid rule, until all rules have been assigned.
17. Brute-force on 3D and 4D grids of booleans. Grids are hash maps to allow negative indexes.
18. Convert infix expressions to postfix with the shunting yard algorithm, then evaluate postfix expressions with the given precedence rules.
19. Convert the given grammar to Chomsky normal form, then use the Cocke–Younger–Kasami algorithm to check if each input line is in the language.
20. -
21. Group foods by allergen, then take set intersect over all ingredient sets inside each group (part A). Recursively search for first valid mapping from allergens to ingredients, where valid means all mapped values are arrays of a single ingredient and each ingredient appears only once in the mapping (part B).
22. Simple recursive search. Copy all cards on each recursive call. Store all seen card permutations in a set of strings.
23. Create a doubly linked circular list and perform 10 000 000 times the un-linking and re-linking steps as defined by the rules.
24. Represent hexagonal grid in axial coordinates and store tile state in a hash table of booleans (true == black tile).
25. Brute-force search.
