function XaydSort(arr)
    if #arr <= 1 then
        return
    end
    xaydSortRecursive(arr, 1, #arr)
end

function xaydSortRecursive(arr, low, high)
    if high - low < 40 then
        xaydSortInsert(arr, low, high)
        return
    end
    
    local pivotIdx = xaydPartition(arr, low, high)
    
    xaydSortRecursive(arr, low, pivotIdx - 1)
    xaydSortRecursive(arr, pivotIdx + 1, high)
end

function xaydPartition(arr, low, high)
    local mid = low + ((high - low) >> 1)
    
    if arr[mid] < arr[low] then
        arr[low], arr[mid] = arr[mid], arr[low]
    end
    if arr[high] < arr[low] then
        arr[low], arr[high] = arr[high], arr[low]
    end
    if arr[mid] < arr[high] then
        arr[mid], arr[high] = arr[high], arr[mid]
    end
    
    local pivot = arr[high]
    local i = low - 1
    
    for j = low, high - 1 do
        if arr[j] <= pivot then
            i = i + 1
            arr[i], arr[j] = arr[j], arr[i]
        end
    end
    
    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return i + 1
end

function xaydSortInsert(arr, low, high)
    for i = low + 1, high do
        local key = arr[i]
        
        local pos = exponentialSearch(arr, low, i - 1, key)
        
        if pos ~= i then
            for k = i, pos + 1, -1 do
                arr[k] = arr[k - 1]
            end
            arr[pos] = key
        end
    end
end

function exponentialSearch(arr, left, right, key)
    if right < left or key <= arr[left] then
        return left
    end
    if key > arr[right] then
        return right + 1
    end
    
    local bound = 1
    while left + bound <= right and arr[left + bound] < key do
        bound = bound << 1
    end
    
    local start = left + (bound >> 1)
    local ending = right
    if left + bound <= right then
        ending = left + bound
    end
    
    return binarySearchInsertPos(arr, start, ending, key)
end

function binarySearchInsertPos(arr, left, right, key)
    while left <= right do
        local mid = left + ((right - left) >> 1)
        
        if arr[mid] < key then
            left = mid + 1
        else
            right = mid - 1
        end
    end
    return left
end
