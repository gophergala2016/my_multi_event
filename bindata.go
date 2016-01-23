// Code generated by go-bindata.
// sources:
// assets/css/materialize.css
// assets/css/materialize.min.css
// assets/font/material-design-icons/LICENSE.txt
// assets/font/material-design-icons/Material-Design-Icons.eot
// assets/font/material-design-icons/Material-Design-Icons.svg
// assets/font/material-design-icons/Material-Design-Icons.ttf
// assets/font/material-design-icons/Material-Design-Icons.woff
// assets/font/material-design-icons/Material-Design-Icons.woff2
// assets/font/roboto/Roboto-Bold.eot
// assets/font/roboto/Roboto-Bold.ttf
// assets/font/roboto/Roboto-Bold.woff
// assets/font/roboto/Roboto-Bold.woff2
// assets/font/roboto/Roboto-Light.eot
// assets/font/roboto/Roboto-Light.ttf
// assets/font/roboto/Roboto-Light.woff
// assets/font/roboto/Roboto-Light.woff2
// assets/font/roboto/Roboto-Medium.eot
// assets/font/roboto/Roboto-Medium.ttf
// assets/font/roboto/Roboto-Medium.woff
// assets/font/roboto/Roboto-Medium.woff2
// assets/font/roboto/Roboto-Regular.eot
// assets/font/roboto/Roboto-Regular.ttf
// assets/font/roboto/Roboto-Regular.woff
// assets/font/roboto/Roboto-Regular.woff2
// assets/font/roboto/Roboto-Thin.eot
// assets/font/roboto/Roboto-Thin.ttf
// assets/font/roboto/Roboto-Thin.woff
// assets/font/roboto/Roboto-Thin.woff2
// assets/js/angular-materialize.js
// assets/js/angular.min.js
// assets/js/jquery-2.2.0.min.js
// assets/js/materialize.js
// assets/js/materialize.min.js
// assets/tmpl/index.html
// DO NOT EDIT!

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// assetsCssMaterializeCss reads file data from disk. It returns an error on failure.
func assetsCssMaterializeCss() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/css/materialize.css"
	name := "assets/css/materialize.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssMaterializeMinCss reads file data from disk. It returns an error on failure.
func assetsCssMaterializeMinCss() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/css/materialize.min.css"
	name := "assets/css/materialize.min.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontMaterialDesignIconsLicenseTxt reads file data from disk. It returns an error on failure.
func assetsFontMaterialDesignIconsLicenseTxt() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/material-design-icons/LICENSE.txt"
	name := "assets/font/material-design-icons/LICENSE.txt"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontMaterialDesignIconsMaterialDesignIconsEot reads file data from disk. It returns an error on failure.
func assetsFontMaterialDesignIconsMaterialDesignIconsEot() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/material-design-icons/Material-Design-Icons.eot"
	name := "assets/font/material-design-icons/Material-Design-Icons.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontMaterialDesignIconsMaterialDesignIconsSvg reads file data from disk. It returns an error on failure.
func assetsFontMaterialDesignIconsMaterialDesignIconsSvg() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/material-design-icons/Material-Design-Icons.svg"
	name := "assets/font/material-design-icons/Material-Design-Icons.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontMaterialDesignIconsMaterialDesignIconsTtf reads file data from disk. It returns an error on failure.
func assetsFontMaterialDesignIconsMaterialDesignIconsTtf() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/material-design-icons/Material-Design-Icons.ttf"
	name := "assets/font/material-design-icons/Material-Design-Icons.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontMaterialDesignIconsMaterialDesignIconsWoff reads file data from disk. It returns an error on failure.
func assetsFontMaterialDesignIconsMaterialDesignIconsWoff() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/material-design-icons/Material-Design-Icons.woff"
	name := "assets/font/material-design-icons/Material-Design-Icons.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontMaterialDesignIconsMaterialDesignIconsWoff2 reads file data from disk. It returns an error on failure.
func assetsFontMaterialDesignIconsMaterialDesignIconsWoff2() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/material-design-icons/Material-Design-Icons.woff2"
	name := "assets/font/material-design-icons/Material-Design-Icons.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoBoldEot reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoBoldEot() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Bold.eot"
	name := "assets/font/roboto/Roboto-Bold.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoBoldTtf reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoBoldTtf() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Bold.ttf"
	name := "assets/font/roboto/Roboto-Bold.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoBoldWoff reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoBoldWoff() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Bold.woff"
	name := "assets/font/roboto/Roboto-Bold.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoBoldWoff2 reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoBoldWoff2() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Bold.woff2"
	name := "assets/font/roboto/Roboto-Bold.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoLightEot reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoLightEot() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Light.eot"
	name := "assets/font/roboto/Roboto-Light.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoLightTtf reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoLightTtf() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Light.ttf"
	name := "assets/font/roboto/Roboto-Light.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoLightWoff reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoLightWoff() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Light.woff"
	name := "assets/font/roboto/Roboto-Light.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoLightWoff2 reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoLightWoff2() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Light.woff2"
	name := "assets/font/roboto/Roboto-Light.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoMediumEot reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoMediumEot() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Medium.eot"
	name := "assets/font/roboto/Roboto-Medium.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoMediumTtf reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoMediumTtf() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Medium.ttf"
	name := "assets/font/roboto/Roboto-Medium.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoMediumWoff reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoMediumWoff() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Medium.woff"
	name := "assets/font/roboto/Roboto-Medium.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoMediumWoff2 reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoMediumWoff2() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Medium.woff2"
	name := "assets/font/roboto/Roboto-Medium.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoRegularEot reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoRegularEot() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Regular.eot"
	name := "assets/font/roboto/Roboto-Regular.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoRegularTtf reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoRegularTtf() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Regular.ttf"
	name := "assets/font/roboto/Roboto-Regular.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoRegularWoff reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoRegularWoff() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Regular.woff"
	name := "assets/font/roboto/Roboto-Regular.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoRegularWoff2 reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoRegularWoff2() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Regular.woff2"
	name := "assets/font/roboto/Roboto-Regular.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoThinEot reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoThinEot() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Thin.eot"
	name := "assets/font/roboto/Roboto-Thin.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoThinTtf reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoThinTtf() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Thin.ttf"
	name := "assets/font/roboto/Roboto-Thin.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoThinWoff reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoThinWoff() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Thin.woff"
	name := "assets/font/roboto/Roboto-Thin.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFontRobotoRobotoThinWoff2 reads file data from disk. It returns an error on failure.
func assetsFontRobotoRobotoThinWoff2() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/font/roboto/Roboto-Thin.woff2"
	name := "assets/font/roboto/Roboto-Thin.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsAngularMaterializeJs reads file data from disk. It returns an error on failure.
func assetsJsAngularMaterializeJs() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/js/angular-materialize.js"
	name := "assets/js/angular-materialize.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsAngularMinJs reads file data from disk. It returns an error on failure.
func assetsJsAngularMinJs() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/js/angular.min.js"
	name := "assets/js/angular.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsJquery220MinJs reads file data from disk. It returns an error on failure.
func assetsJsJquery220MinJs() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/js/jquery-2.2.0.min.js"
	name := "assets/js/jquery-2.2.0.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsMaterializeJs reads file data from disk. It returns an error on failure.
func assetsJsMaterializeJs() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/js/materialize.js"
	name := "assets/js/materialize.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsMaterializeMinJs reads file data from disk. It returns an error on failure.
func assetsJsMaterializeMinJs() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/js/materialize.min.js"
	name := "assets/js/materialize.min.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsTmplIndexHtml reads file data from disk. It returns an error on failure.
func assetsTmplIndexHtml() (*asset, error) {
	path := "/home/andrew/go/src/github.com/acsellers/my_multi_event/assets/tmpl/index.html"
	name := "assets/tmpl/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/css/materialize.css": assetsCssMaterializeCss,
	"assets/css/materialize.min.css": assetsCssMaterializeMinCss,
	"assets/font/material-design-icons/LICENSE.txt": assetsFontMaterialDesignIconsLicenseTxt,
	"assets/font/material-design-icons/Material-Design-Icons.eot": assetsFontMaterialDesignIconsMaterialDesignIconsEot,
	"assets/font/material-design-icons/Material-Design-Icons.svg": assetsFontMaterialDesignIconsMaterialDesignIconsSvg,
	"assets/font/material-design-icons/Material-Design-Icons.ttf": assetsFontMaterialDesignIconsMaterialDesignIconsTtf,
	"assets/font/material-design-icons/Material-Design-Icons.woff": assetsFontMaterialDesignIconsMaterialDesignIconsWoff,
	"assets/font/material-design-icons/Material-Design-Icons.woff2": assetsFontMaterialDesignIconsMaterialDesignIconsWoff2,
	"assets/font/roboto/Roboto-Bold.eot": assetsFontRobotoRobotoBoldEot,
	"assets/font/roboto/Roboto-Bold.ttf": assetsFontRobotoRobotoBoldTtf,
	"assets/font/roboto/Roboto-Bold.woff": assetsFontRobotoRobotoBoldWoff,
	"assets/font/roboto/Roboto-Bold.woff2": assetsFontRobotoRobotoBoldWoff2,
	"assets/font/roboto/Roboto-Light.eot": assetsFontRobotoRobotoLightEot,
	"assets/font/roboto/Roboto-Light.ttf": assetsFontRobotoRobotoLightTtf,
	"assets/font/roboto/Roboto-Light.woff": assetsFontRobotoRobotoLightWoff,
	"assets/font/roboto/Roboto-Light.woff2": assetsFontRobotoRobotoLightWoff2,
	"assets/font/roboto/Roboto-Medium.eot": assetsFontRobotoRobotoMediumEot,
	"assets/font/roboto/Roboto-Medium.ttf": assetsFontRobotoRobotoMediumTtf,
	"assets/font/roboto/Roboto-Medium.woff": assetsFontRobotoRobotoMediumWoff,
	"assets/font/roboto/Roboto-Medium.woff2": assetsFontRobotoRobotoMediumWoff2,
	"assets/font/roboto/Roboto-Regular.eot": assetsFontRobotoRobotoRegularEot,
	"assets/font/roboto/Roboto-Regular.ttf": assetsFontRobotoRobotoRegularTtf,
	"assets/font/roboto/Roboto-Regular.woff": assetsFontRobotoRobotoRegularWoff,
	"assets/font/roboto/Roboto-Regular.woff2": assetsFontRobotoRobotoRegularWoff2,
	"assets/font/roboto/Roboto-Thin.eot": assetsFontRobotoRobotoThinEot,
	"assets/font/roboto/Roboto-Thin.ttf": assetsFontRobotoRobotoThinTtf,
	"assets/font/roboto/Roboto-Thin.woff": assetsFontRobotoRobotoThinWoff,
	"assets/font/roboto/Roboto-Thin.woff2": assetsFontRobotoRobotoThinWoff2,
	"assets/js/angular-materialize.js": assetsJsAngularMaterializeJs,
	"assets/js/angular.min.js": assetsJsAngularMinJs,
	"assets/js/jquery-2.2.0.min.js": assetsJsJquery220MinJs,
	"assets/js/materialize.js": assetsJsMaterializeJs,
	"assets/js/materialize.min.js": assetsJsMaterializeMinJs,
	"assets/tmpl/index.html": assetsTmplIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"materialize.css": &bintree{assetsCssMaterializeCss, map[string]*bintree{}},
			"materialize.min.css": &bintree{assetsCssMaterializeMinCss, map[string]*bintree{}},
		}},
		"font": &bintree{nil, map[string]*bintree{
			"material-design-icons": &bintree{nil, map[string]*bintree{
				"LICENSE.txt": &bintree{assetsFontMaterialDesignIconsLicenseTxt, map[string]*bintree{}},
				"Material-Design-Icons.eot": &bintree{assetsFontMaterialDesignIconsMaterialDesignIconsEot, map[string]*bintree{}},
				"Material-Design-Icons.svg": &bintree{assetsFontMaterialDesignIconsMaterialDesignIconsSvg, map[string]*bintree{}},
				"Material-Design-Icons.ttf": &bintree{assetsFontMaterialDesignIconsMaterialDesignIconsTtf, map[string]*bintree{}},
				"Material-Design-Icons.woff": &bintree{assetsFontMaterialDesignIconsMaterialDesignIconsWoff, map[string]*bintree{}},
				"Material-Design-Icons.woff2": &bintree{assetsFontMaterialDesignIconsMaterialDesignIconsWoff2, map[string]*bintree{}},
			}},
			"roboto": &bintree{nil, map[string]*bintree{
				"Roboto-Bold.eot": &bintree{assetsFontRobotoRobotoBoldEot, map[string]*bintree{}},
				"Roboto-Bold.ttf": &bintree{assetsFontRobotoRobotoBoldTtf, map[string]*bintree{}},
				"Roboto-Bold.woff": &bintree{assetsFontRobotoRobotoBoldWoff, map[string]*bintree{}},
				"Roboto-Bold.woff2": &bintree{assetsFontRobotoRobotoBoldWoff2, map[string]*bintree{}},
				"Roboto-Light.eot": &bintree{assetsFontRobotoRobotoLightEot, map[string]*bintree{}},
				"Roboto-Light.ttf": &bintree{assetsFontRobotoRobotoLightTtf, map[string]*bintree{}},
				"Roboto-Light.woff": &bintree{assetsFontRobotoRobotoLightWoff, map[string]*bintree{}},
				"Roboto-Light.woff2": &bintree{assetsFontRobotoRobotoLightWoff2, map[string]*bintree{}},
				"Roboto-Medium.eot": &bintree{assetsFontRobotoRobotoMediumEot, map[string]*bintree{}},
				"Roboto-Medium.ttf": &bintree{assetsFontRobotoRobotoMediumTtf, map[string]*bintree{}},
				"Roboto-Medium.woff": &bintree{assetsFontRobotoRobotoMediumWoff, map[string]*bintree{}},
				"Roboto-Medium.woff2": &bintree{assetsFontRobotoRobotoMediumWoff2, map[string]*bintree{}},
				"Roboto-Regular.eot": &bintree{assetsFontRobotoRobotoRegularEot, map[string]*bintree{}},
				"Roboto-Regular.ttf": &bintree{assetsFontRobotoRobotoRegularTtf, map[string]*bintree{}},
				"Roboto-Regular.woff": &bintree{assetsFontRobotoRobotoRegularWoff, map[string]*bintree{}},
				"Roboto-Regular.woff2": &bintree{assetsFontRobotoRobotoRegularWoff2, map[string]*bintree{}},
				"Roboto-Thin.eot": &bintree{assetsFontRobotoRobotoThinEot, map[string]*bintree{}},
				"Roboto-Thin.ttf": &bintree{assetsFontRobotoRobotoThinTtf, map[string]*bintree{}},
				"Roboto-Thin.woff": &bintree{assetsFontRobotoRobotoThinWoff, map[string]*bintree{}},
				"Roboto-Thin.woff2": &bintree{assetsFontRobotoRobotoThinWoff2, map[string]*bintree{}},
			}},
		}},
		"js": &bintree{nil, map[string]*bintree{
			"angular-materialize.js": &bintree{assetsJsAngularMaterializeJs, map[string]*bintree{}},
			"angular.min.js": &bintree{assetsJsAngularMinJs, map[string]*bintree{}},
			"jquery-2.2.0.min.js": &bintree{assetsJsJquery220MinJs, map[string]*bintree{}},
			"materialize.js": &bintree{assetsJsMaterializeJs, map[string]*bintree{}},
			"materialize.min.js": &bintree{assetsJsMaterializeMinJs, map[string]*bintree{}},
		}},
		"tmpl": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{assetsTmplIndexHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
