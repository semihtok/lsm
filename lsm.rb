# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Lsm < Formula
  desc "A high-performance, concurrent file system analysis tool that efficiently calculates and displays detailed directory and file sizes, optimized for handling large directories and complex file structures."
  homepage "https://github.com/semihtok/lsm"
  version "0.1.2"
  license "Apache 2.0"

  depends_on "git"
  depends_on "go"

  on_macos do
    url "https://github.com/semihtok/lsm/releases/download/v0.1.2/lsm_0.1.2_darwin_amd64.tar.gz"
    sha256 "30f797846626f7419a5fe82eea304190c603008630b546ac9fe3c1709294e4ac"

    def install
      bin.install "lsm"
    end

    if Hardware::CPU.arm?
      def caveats
        <<~EOS
          The darwin_arm64 architecture is not supported for the lsm
          formula at this time. The darwin_amd64 binary may work in compatibility
          mode, but it might not be fully supported.
        EOS
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/semihtok/lsm/releases/download/v0.1.2/lsm_0.1.2_linux_armv6.tar.gz"
      sha256 "839c0451605c90ad0a221384cd7cb797990a35305620a375fcb42628ce51fdb4"

      def install
        bin.install "lsm"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/semihtok/lsm/releases/download/v0.1.2/lsm_0.1.2_linux_amd64.tar.gz"
      sha256 "2fba15bace295e3a62070921762a8f05bc417695864e47cedd82a744bf10894b"

      def install
        bin.install "lsm"
      end
    end
  end
end
