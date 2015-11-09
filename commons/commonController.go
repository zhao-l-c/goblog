package commons

func Diff(new, old []string) (insertSlice, deleteSlice []string) {
    length := len(new)
    oldLen := len(old)
    if length < oldLen {
        length = oldLen
    }
    insert := make([]string, 0, length)
    delete := make([]string, 0, length)
    unchange := make([]string, 0, length)
    for _, a := range new {
        for _, b := range old {
            if a == b {
                unchange = append(unchange, a)
                break
            }
        }
    }
    for _, ele := range new {
        if !Contains(ele, unchange) {
            insert = append(insert, ele)
        }
    }
    for _, ele := range old {
        if !Contains(ele, unchange) {
            delete = append(delete, ele)
        }
    }
    return insert, delete
}

func Contains(s string, slice []string) bool {
    for _, ele := range slice {
        if ele == s {
            return true
        }
    }
    return false
}