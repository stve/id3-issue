require 'rubygems'
require 'bundler/setup'

require 'fileutils'
require 'taglib'

# now that files have been converted to mp3's, update the ID3 tags
puts 'Detecting MP3 files...'
Dir.glob('*.mp3') do |file|
  puts '*' * 20
  puts "File: #{File.basename(file)}"

    TagLib::MPEG::File.open(file) do |mp3_file|
      tag = mp3_file.id3v2_tag
      tag.frame_list.each do |f|
        puts f.frame_id
        puts f.description if f.respond_to?(:description)
        puts f.to_string
        if f.frame_id == 'APIC'
          puts f.description
          puts f.mime_type
          puts f.text_encoding
          puts f.type
        end
      end
      puts
    end
end
