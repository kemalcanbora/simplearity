package singularity

import (
	"fmt"
)

func GenerateScript(jobName, mem, partition string, cpus int, dockerImage string) string {
	script := `#!/bin/bash
#SBATCH --job-name=%s
#SBATCH --cpus-per-task=%d
#SBATCH --mem=%s
#SBATCH --partition=%s
#SBATCH -o %%x-%%j.out          # File to which STDOUT will be written
#SBATCH -e %%x-%%j.err          # File to which STDERR will be written


# Now, build your container
singularity build %s_container.sif docker://%s

echo "Job submitted successfully"
`

	return fmt.Sprintf(script, jobName, cpus, mem, partition, jobName, dockerImage)
}
