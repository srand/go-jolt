package jolt

type depsInfluence struct {
	job Job
}

func DepsInfluence(job Job) Influence {
	return &depsInfluence{job: job}
}

func (d *depsInfluence) Digest() (Digest, error) {
	a := AggregatedInfluence{}

	for _, dep := range d.job.Deps() {
		digest, err := GetXattrDigest(string(dep), XattrJobDigest)
		if err != nil {
			return "", err
		}

		a.AddInfluence(StringInfluence(string(digest)))
	}

	return a.Digest()
}
