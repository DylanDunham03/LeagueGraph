export function calculateLocalStorageSize() {
    let total = 0;
    for (let key in localStorage) {
        if (localStorage.hasOwnProperty(key)) {
            total += localStorage[key].length * 2; // Each character in a string is 16 bits = 2 bytes
        }
    }
    return total / 1024; // Convert bytes to kilobytes
}
