package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GA4H7kyiAEJiDBwMhSXJ6anpxbpQ2m9rOL8vNAQVgZG8SPTE7aE7s+/7CBy73d4BDuzc+fdDanJqTqvWHjCfx/dG3Q/9v3cUNcTzDwrQp62u4hnsIYEXNRd5fDK4Z+sdeKp/Q3n/KTmbF7rwcY/8d3d8vvV32vfx85ba9n3WMBmQSJH1xeZS28rUvZ2GN0TX/JcZIbNI+51K9f7GpxcfJJHa+W9g0bXegWndOs82ajzVoSdJUs76o2G4BE7iW3u1/L6FgS9M5yfpbZjd3eTMPe3lLvPrprkPJCZH5PWfv3mzddfXtSY2vUs7ljVeDGn7Mfx1F+fLKqLJ+7766XF57Eqg0UqpFZym8znfpaGfUmXYm4t0Znk23b95tXurhTz4xa9Bx40Zi9j9/WY9jmve4Hg0gT7NC6mS8VbuhYydv7jc9ly6bHQu21N+07WnOZdcN7C2NzFXIjxTek2gSu1J+U2Tcg/6Xtj+9yZ9Z5P+2TXFgUxeh0U3LAt4tLXPcdmR3zb88vO0l7uf9X8WqvsSZvnenw+ZM0f8OHszUV73Z6eSAqV+7mfS5O/rrf0mvACdqN83417uffFBZ20YRaakma6TuDcn+Wbn637mHOuP8pQ7tYmlanGXRqtq+pmzTFbw+7E+r996+2mbedeuloF3U35f8M3noXXT4/3uuz66bL7vPvt/N2XS5dfzg+fPF/88sWXj298fnfwWMmXj99Ptp8ssPj3aU/al7yu1Jjmn/ejPd83vPrlGbvvv1ePoMKlDb7c139uD/ujVs/EwPD/f4A3O8eWpn+7eJkYGB6wMzDA0hUDwwW0dMWOSFfgpCRxZHoCSDeymgBvRiYRZkS6RDYZlC5h4H8jiMSbShFGYXcKBAgw/HeMYGLA4jBWNpA8EwMTQycDA8M0JhAPEAAA//9UJTp/NQMAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
