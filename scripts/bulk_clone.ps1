[cmdletbinding()]
param(
    [parameter(
        Mandatory         = $true,
        ValueFromPipeline = $true)]
    $pipelineInput
)
	Begin {
		$OutPath = "output"
	}

    Process {

        ForEach ($input in $pipelineInput) {
           $name=$($("$input" -split "/")[-1] -split "\.")[-2]
           git clone --mirror $input $OutPath"\\""raw""\\"$name
           git clone -l $OutPath"\\""raw""\\"$name $OutPath"\\""repo""\\"$name
        }

    }
