package main

import "bunster-build/runtime"

func Main(shell *runtime.Shell, streamManager *runtime.StreamManager) {
	shell.ExitCode = 0
	expr3, exitCode := func() (string, int) {
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		shell := shell.Clone()
		var buffer = runtime.NewBuffer(``, false)
		streamManager.Add(`1`, buffer, false)
		func() {
			var commandName = `xdg-user-dir`
			var arguments []string
			arguments = append(arguments, `PICTURES`)
			var command = shell.Command(commandName, arguments...)
			streamManager := streamManager.Clone()
			defer streamManager.Destroy()
			if stream, err := streamManager.Get(`0`); err != nil {
				shell.HandleError(err)
				return
			} else {
				command.Stdin = stream
			}
			if stream, err := streamManager.Get(`1`); err != nil {
				shell.HandleError(err)
				return
			} else {
				command.Stdout = stream
			}
			if stream, err := streamManager.Get(`2`); err != nil {
				shell.HandleError(err)
				return
			} else {
				command.Stderr = stream
			}
			if err := command.Run(); err != nil {
				shell.HandleError(err)
				return
			}
			shell.ExitCode = command.ProcessState.ExitCode()

		}()
		return buffer.String(true), shell.ExitCode
	}()
	if exitCode != 0 {
		shell.ExitCode = exitCode
		return
	}
	shell.SetVar("BASE_PICTURES_FOLDER", expr3)
	shell.ExitCode = 0
	shell.SetVar("PICTURES_FOLDER", shell.ReadVar("BASE_PICTURES_FOLDER")+`/xkcd_strips`)
	func() {
		var commandName = `mkdir`
		var arguments []string
		arguments = append(arguments, shell.ReadVar("PICTURES_FOLDER"))
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		stream0, err := streamManager.OpenStream(`/dev/null`, runtime.STREAM_FLAG_WRITE)
		if err != nil {
			shell.HandleError(err)
			return
		}
		streamManager.Add(`2`, stream0, false)
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	func() {
		var commandName = `echo`
		var arguments []string
		arguments = append(arguments, `[+] Downloading random xkcd comic...`)
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	func() {
		var commandName = `wget`
		var arguments []string
		arguments = append(arguments, `-q`)
		arguments = append(arguments, `https://c.xkcd.com/random/comic/`)
		arguments = append(arguments, `-O`)
		arguments = append(arguments, `/tmp/comics_index.html`)
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	func() {
		var commandName = `echo`
		var arguments []string
		arguments = append(arguments, `[*] Parsing the comic...`)
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	shell.ExitCode = 0
	expr24, exitCode := func() (string, int) {
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		shell := shell.Clone()
		var buffer = runtime.NewBuffer(``, false)
		streamManager.Add(`1`, buffer, false)
		func() {
			var pipelineWaitgroup []func() error
			pipeReader1, pipeWriter1, err := runtime.NewPipe()
			if err != nil {
				shell.HandleError(err)
				return
			}
			func() {
				var commandName = `grep`
				var arguments []string
				arguments = append(arguments, `og:title`)
				arguments = append(arguments, `/tmp/comics_index.html`)
				var command = shell.Command(commandName, arguments...)
				streamManager := streamManager.Clone()
				streamManager.Add(`1`, pipeWriter1, true)
				if stream, err := streamManager.Get(`0`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stdin = stream
				}
				if stream, err := streamManager.Get(`1`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stdout = stream
				}
				if stream, err := streamManager.Get(`2`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stderr = stream
				}
				if err := command.Start(); err != nil {
					shell.HandleError(err)
					return
				}
				pipelineWaitgroup = append(pipelineWaitgroup, func() error {
					defer streamManager.Destroy()
					return command.Wait()
				})

			}()
			func() {
				var commandName = `sed`
				var arguments []string
				arguments = append(arguments, `s/.*content="\([^"]*\)".*/\1/`)
				var command = shell.Command(commandName, arguments...)
				streamManager := streamManager.Clone()
				streamManager.Add(`0`, pipeReader1, false)
				if stream, err := streamManager.Get(`0`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stdin = stream
				}
				if stream, err := streamManager.Get(`1`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stdout = stream
				}
				if stream, err := streamManager.Get(`2`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stderr = stream
				}
				if err := command.Start(); err != nil {
					shell.HandleError(err)
					return
				}
				pipelineWaitgroup = append(pipelineWaitgroup, func() error {
					defer streamManager.Destroy()
					return command.Wait()
				})

			}()
			for i, wait := range pipelineWaitgroup {
				if err := wait(); err != nil {
					shell.HandleError(err)
				}
				if i < (len(pipelineWaitgroup) - 1) {
					shell.ExitCode = 0
				}
			}

		}()
		return buffer.String(true), shell.ExitCode
	}()
	if exitCode != 0 {
		shell.ExitCode = exitCode
		return
	}
	shell.SetVar("title", expr24)
	shell.ExitCode = 0
	expr30, exitCode := func() (string, int) {
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		shell := shell.Clone()
		var buffer = runtime.NewBuffer(``, false)
		streamManager.Add(`1`, buffer, false)
		func() {
			var pipelineWaitgroup []func() error
			pipeReader1, pipeWriter1, err := runtime.NewPipe()
			if err != nil {
				shell.HandleError(err)
				return
			}
			func() {
				var commandName = `grep`
				var arguments []string
				arguments = append(arguments, `og:image`)
				arguments = append(arguments, `/tmp/comics_index.html`)
				var command = shell.Command(commandName, arguments...)
				streamManager := streamManager.Clone()
				streamManager.Add(`1`, pipeWriter1, true)
				if stream, err := streamManager.Get(`0`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stdin = stream
				}
				if stream, err := streamManager.Get(`1`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stdout = stream
				}
				if stream, err := streamManager.Get(`2`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stderr = stream
				}
				if err := command.Start(); err != nil {
					shell.HandleError(err)
					return
				}
				pipelineWaitgroup = append(pipelineWaitgroup, func() error {
					defer streamManager.Destroy()
					return command.Wait()
				})

			}()
			func() {
				var commandName = `sed`
				var arguments []string
				arguments = append(arguments, `s/.*content="\([^"]*\)".*/\1/`)
				var command = shell.Command(commandName, arguments...)
				streamManager := streamManager.Clone()
				streamManager.Add(`0`, pipeReader1, false)
				if stream, err := streamManager.Get(`0`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stdin = stream
				}
				if stream, err := streamManager.Get(`1`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stdout = stream
				}
				if stream, err := streamManager.Get(`2`); err != nil {
					shell.HandleError(err)
					return
				} else {
					command.Stderr = stream
				}
				if err := command.Start(); err != nil {
					shell.HandleError(err)
					return
				}
				pipelineWaitgroup = append(pipelineWaitgroup, func() error {
					defer streamManager.Destroy()
					return command.Wait()
				})

			}()
			for i, wait := range pipelineWaitgroup {
				if err := wait(); err != nil {
					shell.HandleError(err)
				}
				if i < (len(pipelineWaitgroup) - 1) {
					shell.ExitCode = 0
				}
			}

		}()
		return buffer.String(true), shell.ExitCode
	}()
	if exitCode != 0 {
		shell.ExitCode = exitCode
		return
	}
	shell.SetVar("imageurl", expr30)
	func() {
		var commandName = `echo`
		var arguments []string
		arguments = append(arguments, shell.ReadVar("title"))
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	func() {
		var commandName = `echo`
		var arguments []string
		arguments = append(arguments, shell.ReadVar("imageurl"))
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	func() {
		var commandName = `echo`
		var arguments []string
		arguments = append(arguments, `[+] Downloading the image...`)
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	func() {
		var commandName = `curl`
		var arguments []string
		arguments = append(arguments, `-s`)
		arguments = append(arguments, shell.ReadVar("imageurl"))
		arguments = append(arguments, `-o`)
		arguments = append(arguments, shell.ReadVar("PICTURES_FOLDER")+`/`+shell.ReadVar("title")+`.png`)
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	func() {
		var commandName = `echo`
		var arguments []string
		arguments = append(arguments, `[+] Saving the image to a file...`)
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	func() {
		var commandName = `echo`
		var arguments []string
		arguments = append(arguments, `[Ok] Done (`+shell.ReadVar("PICTURES_FOLDER")+`/`+shell.ReadVar("title")+`.png)!`)
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	func() {
		var commandName = `rm`
		var arguments []string
		arguments = append(arguments, `-rf`)
		arguments = append(arguments, `/tmp/comics_index.html`)
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
	func() {
		var commandName = `kitten`
		var arguments []string
		arguments = append(arguments, `icat`)
		arguments = append(arguments, shell.ReadVar("PICTURES_FOLDER")+`/`+shell.ReadVar("title")+`.png`)
		var command = shell.Command(commandName, arguments...)
		streamManager := streamManager.Clone()
		defer streamManager.Destroy()
		stream0, err := streamManager.OpenStream(`/dev/null`, runtime.STREAM_FLAG_WRITE)
		if err != nil {
			shell.HandleError(err)
			return
		}
		streamManager.Add(`2`, stream0, false)
		if stream, err := streamManager.Get(`0`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdin = stream
		}
		if stream, err := streamManager.Get(`1`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stdout = stream
		}
		if stream, err := streamManager.Get(`2`); err != nil {
			shell.HandleError(err)
			return
		} else {
			command.Stderr = stream
		}
		if err := command.Run(); err != nil {
			shell.HandleError(err)
			return
		}
		shell.ExitCode = command.ProcessState.ExitCode()

	}()
}
