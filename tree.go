package treewriter

import (
	"fmt"
	"io"
)

type Tree interface {
	Children() []Tree
}

func WriteTo(w io.Writer, t Tree, conf *Config) (int64, error) {
	if conf == nil {
		conf = DefaultConfig
	}

	return write(w, t, conf, "")
}

func write(writer io.Writer, tree Tree, conf *Config, prefix string) (int64, error) {
	n, err := fmt.Fprintf(writer, "%s%v%s", prefix, tree, conf.newline)
	if err != nil {
		return int64(n), err
	}

	written := int64(n)
	childs := tree.Children()
	childsLen := len(childs)
	prefix = conf.replacer.Replace(prefix)

	for i, tree := range childs {
		var prefixAdd string
		if i == childsLen-1 && len(tree.Children()) == 0 {
			prefixAdd = conf.endPrefix
		} else {
			prefixAdd = conf.midPrefix
		}

		nt, err := write(writer, tree, conf, prefix+prefixAdd)
		if err != nil {
			return written + nt, err
		}

		written += nt
	}

	return written, nil
}
