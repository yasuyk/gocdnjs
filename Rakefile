PROJECT_ROOT_DIR = File.dirname __FILE__
VERSION_FILE = File.join PROJECT_ROOT_DIR , 'version.go'
VERSION_REGEX = Regexp.compile '"(\d+\.\d+\.\d+)"'

OPERATION_SYSEMS = %W/windows linux darwin/
ARCHITECTURES = %W/amd64 386/
APPNAME = 'gocdnjs'

def version
  version = ''
  File.readlines(VERSION_FILE).each do |l|
    data = VERSION_REGEX.match(l)
    if data
      version = data[1]
      break
    end
  end
  version
end

def exename os
  if os == 'windows'
    "#{APPNAME}.exe"
  else
    APPNAME
  end
end

def say(text, color = :magenta)
  n = { bold: 1, red: 31, green: 32, yellow: 33, blue: 34, magenta: 35 }.fetch(color, 0)
  puts "\e[%dm%s\e[0m" % [n, text]
end

desc 'Release executable'
task :release do
  dir = File.join PROJECT_ROOT_DIR, 'release'
  FileUtils.mkdir_p dir

  OPERATION_SYSEMS.each do |os|
    ARCHITECTURES.each do |arch|
      exe = File.join dir, exename(os)
      sh "GOOS=#{os} GOARCH=#{arch} go build -o #{exe}"  do |ok, res|
        if ok
          zip = File.join dir, "#{APPNAME}_#{version}_#{os}_#{arch}.zip"
          sh "zip #{zip} #{exe}" do |ok, res|
            FileUtils.rm exe
          end
        end
      end
    end
  end
end
