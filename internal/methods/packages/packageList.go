package packages

import "strings"


const (
	NVIDIAPACKAGES = iota // 0 
	DEVELOPMENTPACKAGES  // 1
	// We can add more later...
)

var PackageMap map[int]PackageGroup

func InitPackageMap(){
	PackageMap = make(map[int]PackageGroup)
	PackageMap[NVIDIAPACKAGES] = PackageGroup{
		PackageID: NVIDIAPACKAGES,
		Name: "nvidia",
		Packages: []string{"go", "vim"}, // Placeholders, obviously these are not NVIDIA packages
	}
	PackageMap[DEVELOPMENTPACKAGES] = PackageGroup{
		PackageID: DEVELOPMENTPACKAGES,
		Name: "development",
		Packages: []string{"git", "go"},
	}


}



func DeterminePkg(pkgGrp string) (PackageGroup, error){
	switch strings.ToLower(pkgGrp) {
	case PackageMap[NVIDIAPACKAGES].Name:
		return PackageMap[NVIDIAPACKAGES], nil
	case PackageMap[DEVELOPMENTPACKAGES].Name:
		return PackageMap[DEVELOPMENTPACKAGES], nil
	default:
		return PackageGroup{}, ErrNotFound
	}
}