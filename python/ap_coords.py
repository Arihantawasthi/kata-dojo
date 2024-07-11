# coords = [-25, -18, -12, -7, -3, 0, 3, 7, 12, 18, 25]


def gen_coords(initial_diff, limit):
    initial_arr = [0 for _ in range(0, (limit - 1) + (limit - 2))]
    mid_point = limit - 2
    l_pointer = mid_point - 1
    r_pointer = mid_point
    val = 0

    while (r_pointer < len(initial_arr)):
        if r_pointer == mid_point:
            initial_arr[r_pointer] = val
            r_pointer += 1
            val += initial_diff
            continue

        initial_arr[r_pointer] = val
        initial_arr[l_pointer] = -val
        initial_diff += 1
        val += initial_diff
        r_pointer += 1
        l_pointer -= 1

    return initial_arr


coords = gen_coords(3, 7)
print(coords)
