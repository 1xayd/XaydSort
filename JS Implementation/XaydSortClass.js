//this was written by ai im not a js programmer.
class XaydSort {
    constructor(arr, visualizer = null) {
        this.arr = arr;
        this.visualizer = visualizer;
    }

    async sort() {
        if (this.arr.length <= 1) {
            return;
        }
        await this.xaydSortRecursive(0, this.arr.length - 1);
    }

    async xaydSortRecursive(low, high) {
        if (high - low < 40) {
            await this.xaydSortInsert(low, high);
            return;
        }

        const pivotIdx = await this.xaydPartition(low, high);

        await this.xaydSortRecursive(low, pivotIdx - 1);
        await this.xaydSortRecursive(pivotIdx + 1, high);
    }

    async xaydPartition(low, high) {
        const mid = low + Math.floor((high - low) / 2);

        await this.compare(mid, low);
        if (this.arr[mid] < this.arr[low]) {
            await this.swap(low, mid);
        }
        
        await this.compare(high, low);
        if (this.arr[high] < this.arr[low]) {
            await this.swap(low, high);
        }
        
        await this.compare(mid, high);
        if (this.arr[mid] < this.arr[high]) {
            await this.swap(mid, high);
        }

        const pivot = this.arr[high];
        let i = low - 1;

        for (let j = low; j < high; j++) {
            await this.compare(j, high);
            if (this.arr[j] <= pivot) {
                i++;
                if (i !== j) {
                    await this.swap(i, j);
                }
            }
        }

        await this.swap(i + 1, high);
        return i + 1;
    }

    async xaydSortInsert(low, high) {
        for (let i = low + 1; i <= high; i++) {
            const key = this.arr[i];
            await this.markSpecial(i);

            const pos = await this.exponentialSearch(low, i - 1, key);

            if (pos !== i) {
                for (let k = i; k > pos; k--) {
                    this.arr[k] = this.arr[k - 1];
                    await this.updateVisualizer();
                }
                this.arr[pos] = key;
                await this.updateVisualizer();
            }
        }
    }

    async exponentialSearch(left, right, key) {
        if (right < left || key <= this.arr[left]) {
            return left;
        }
        if (key > this.arr[right]) {
            return right + 1;
        }

        let bound = 1;
        while (left + bound <= right && this.arr[left + bound] < key) {
            await this.compare(left + bound, -1, key);
            bound <<= 1;
        }

        const start = left + (bound >> 1);
        const end = left + bound <= right ? left + bound : right;

        return await this.binarySearchInsertPos(start, end, key);
    }

    async binarySearchInsertPos(left, right, key) {
        while (left <= right) {
            const mid = left + Math.floor((right - left) / 2);
            await this.compare(mid, -1, key);

            if (this.arr[mid] < key) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return left;
    }

    async compare(idx1, idx2, value = null) {
        if (this.visualizer) {
            await this.visualizer.compare(idx1, idx2, value);
        }
    }

    async swap(idx1, idx2) {
        if (this.visualizer) {
            await this.visualizer.swap(idx1, idx2);
        }
        [this.arr[idx1], this.arr[idx2]] = [this.arr[idx2], this.arr[idx1]];
    }

    async markSpecial(idx) {
        if (this.visualizer) {
            await this.visualizer.markSpecial(idx);
        }
    }

    async updateVisualizer() {
        if (this.visualizer) {
            await this.visualizer.update();
        }
    }
}

if (typeof module !== 'undefined' && module.exports) {
    module.exports = XaydSort;
}
