/**
 * @param {number[]} heights
 * @return {number}
 */
var largestRectangleArea = function(heights) {
    let maxArea = 0;
    const stack = []; // Stores indices
    // Add a 0 height at the end to ensure all bars are processed
    const newHeights = [...heights, 0];

    for (let i = 0; i < newHeights.length; i++) {
        // While current bar is shorter than the bar at stack's top
        while (stack.length > 0 && newHeights[i] < newHeights[stack[stack.length - 1]]) {
            const h = newHeights[stack.pop()]; // Height of the rectangle
            // If stack is empty, width spans from 0 to i
            // Else, width spans from the index after the new stack top to i
            const w = stack.length === 0 ? i : i - stack[stack.length - 1] - 1;
            maxArea = Math.max(maxArea, h * w);
        }
        stack.push(i);
    }

    return maxArea;
};
