package packages


const (
	NVIDIAPACKAGES = iota // 0 
	DEVELOPMENTPACKAGES  // 1
	// We can add more later...
)

var PackageMap map[int]PackageGroup

func InitPackageMap(){
	PackageMap = make(map[int]PackageGroup)
	PackageMap[NVIDIAPACKAGES] = PackageGroup{
		Name: "NVIDIA",
		Packages: []string{"go", "vim"}, // Placeholders, obviously these are not NVIDIA packages
	}
	PackageMap[DEVELOPMENTPACKAGES] = PackageGroup{
		Name: "Development",
		Packages: []string{"git", "go"},
	}


}

func DeterminePkg(pkgGrp string) (PackageGroup, error){
	switch pkgGrp {
	case PackageMap[NVIDIAPACKAGES].Name:
		return PackageMap[NVIDIAPACKAGES], nil
	case PackageMap[DEVELOPMENTPACKAGES].Name:
		return PackageMap[DEVELOPMENTPACKAGES], nil
	default:
		return PackageGroup{}, ErrNotFound
	}
}