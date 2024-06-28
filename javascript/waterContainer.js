const heights = [1,8,6,2,5,4,8,3,7]

function maxArea(heights) {
    let pRight = heights.length - 1;
    let pLeft = 0;
    let area = 0;
    let width = 0;
    let tall = 0;

    while (pLeft < pRight) {
        width = pRight - pLeft;
        tall = heights[pLeft];
        if (tall > heights[pRight]) {
            tall = heights[pRight];
            pRight -= 1;
        } else {
            pLeft += 1;
        }

        area = Math.max(area, width * tall);
    }

    console.log(area);
    return area;
}

maxArea(heights)
