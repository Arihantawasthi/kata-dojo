const string1= "accabbcbba";
const string2 = "dfdd";


function longestSubstring(string) {
    let length = string.length;
    let seen = new Set()
    let max = 0
    let lPointer = 0;

    for (let rPointer = 0; rPointer < length; rPointer++) {
        if (seen.has(string[rPointer])) {
            lPointer++;
        }
        seen.add(string[rPointer]);
        max = Math.max(max, (rPointer - lPointer) + 1);
    }

    return max;
}


console.log(longestSubstring(string1));
console.log(longestSubstring(string2));
