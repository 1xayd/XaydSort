//implementation in https://sortvisualizer.com/customsort/#box
//github at: https://github.com/1xayd/XaydSort
async function xaydSort(elements) {
    if (elements.length <= 1) {
        return;
    }
    await xaydSortRecursive(elements, 0, elements.length - 1);
}

async function xaydSortRecursive(elements, low, high) {
    if (high - low < 40) {
        await xaydSortInsert(elements, low, high);
        return;
    }

    const pivotIdx = await xaydPartition(elements, low, high);

    await xaydSortRecursive(elements, low, pivotIdx - 1);
    await xaydSortRecursive(elements, pivotIdx + 1, high);
}

async function xaydPartition(elements, low, high) {
    const mid = low + ((high - low) >> 1);

    // Median-of-three pivot selection
    if (getValue(elements[mid]) < getValue(elements[low])) {
        await swap(low, mid);
    }
    if (getValue(elements[high]) < getValue(elements[low])) {
        await swap(low, high);
    }
    if (getValue(elements[mid]) < getValue(elements[high])) {
        await swap(mid, high);
    }

    const pivot = getValue(elements[high]);
    let i = low - 1;

    for (let j = low; j < high; j++) {
        if (getValue(elements[j]) <= pivot) {
            i++;
            if (i !== j) {
                await swap(i, j);
            }
        }
    }

    await swap(i + 1, high);
    return i + 1;
}

async function xaydSortInsert(elements, low, high) {
    for (let i = low + 1; i <= high; i++) {
        const key = getValue(elements[i]);

        const pos = exponentialSearch(elements, low, i - 1, key);

        if (pos !== i) {
            // Store the element to insert
            const temp = elements[i];

            // Shift elements to the right
            for (let k = i; k > pos; k--) {
                elements[k] = elements[k - 1];
            }

            // Insert the element at the correct position
            elements[pos] = temp;
        }
		
        await updateBox(elements, pos);
    }
}

function exponentialSearch(elements, left, right, key) {
    if (right < left || key <= getValue(elements[left])) {
        return left;
    }
    if (key > getValue(elements[right])) {
        return right + 1;
    }

    let bound = 1;
    while (left + bound <= right && getValue(elements[left + bound]) < key) {
        bound <<= 1;
    }

    const start = left + (bound >> 1);
    let end = right;
    if (left + bound <= right) {
        end = left + bound;
    }

    return binarySearchInsertPos(elements, start, end, key);
}

function binarySearchInsertPos(elements, left, right, key) {
    while (left <= right) {
        const mid = left + ((right - left) >> 1);

        if (getValue(elements[mid]) < key) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return left;
}
