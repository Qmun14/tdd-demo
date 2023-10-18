package twosum

// func TwoSum(nums []int, target int) []int {
// 	firstIdx := 0
// 	maxIdx := len(nums) - 1

// 	for firstIdx < maxIdx {
// 		secondIdx := firstIdx + 1

// 		for secondIdx <= maxIdx {
// 			if nums[firstIdx]+nums[secondIdx] == target {
// 				return []int{firstIdx, secondIdx}
// 			}
// 			secondIdx++
// 		}
// 		firstIdx++
// 	}

// 	return nil
// }

func TwoSum(nums []int, target int) []int {

	// numMap adalah sebuah peta (map) yang digunakan untuk memetakan angka-angka yang telah dilihat sebelumnya dalam slice nums ke indeks mereka dalam slice.
	numMap := make(map[int]int)

	for i, num := range nums {
		// Jika complement ada dalam peta(map) dan ditemukan (artinya ada pasangan angka yang saat dijumlahkan akan sama dengan target), maka found akan bernilai true, dan idx akan berisi indeks dari angka tersebut dalam slice nums.
		complement := target - num

		if idx, found := numMap[complement]; found {
			// Jika kondisi found adalah true, artinya ditemukan pasangan angka yang cocok, maka fungsi mengembalikan sebuah slice dengan dua elemen, yaitu idx (indeks angka dalam slice nums yang saat dijumlahkan dengan complement akan sama dengan target) dan i (indeks angka saat ini dalam iterasi).
			return []int{idx, i}
		}

		numMap[num] = i
	}

	return []int{}
}
