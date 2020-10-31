# This outlines the process for solving questions in an interview

1. Listen carefully, store all the special bits of information that the interviewer gives you as it might be useful in finding an optimal algorithm(e.g if the array is sorted)
2. Example. Use a good example. One that is 
- Not a special cases
- Sufficiently large
- Specific(use real numbers and strings)
3. Brute force: Start with a brute force algorithm, do not implement it yet and state the time and space complexity
4. Optimize: Look for ways to optimize the algorithm
- Are you repeating computations[]
- Are there an unused part of the question?
- Could you make a space/time tradeoff
- Could you use a hashmap or another data structure
- Can you precompute information

Optimization Techniques #1: BUD
- Bottlenecks, Unnecessay work and duplication
Bottlenecks: Check for a part of your algorithm that slows down the overall runtime.

Optimization Techniques #2: DIY
- Do it yourself
Try working through it intuitively with a real example and you will find that you don't use a brute force algorithm but rather an optimized version.

Optimization Techniques #3: Simplify and Generalize
- You can solve the algorithm for a simpler data structure then modify and tweak it for a more complex one.

Optimization Techniques #4: Base Case and build
- You solve the problem for a base case of 1, then try to build up from there. When solving more complex cases such as i = 3, 4. You try and use prior solutions

Optimization Techniques #4: Data Structure BrainStor,
- You move through your list of ds and find one suited to the problem.

5. Walkthrough
- Walkthrough your code carefully so that you get a feel for the structure of the code.
- Understand exactly the code you want to write
6. Implement
- Write modularized code
- Have error checks(You can add a todo so you can comeback to it later.)
- Use classes/structs when possible
- Use meaningful variable names
7. Test
- Test conceptually. Think about what each line of the code does.
- Look out for weird cases like the loop that starts at 1
- Check for hotspots like null values in trees, the end of an array etc
- Use small test cases
- Test for edge cases