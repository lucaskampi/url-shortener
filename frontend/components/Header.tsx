'use client'

import React from 'react'
import Link from 'next/link'
import { Link as LinkIcon } from 'lucide-react'

const Header: React.FC = () => {
  return (
    <header className="fixed top-7 left-0 right-0 z-50">
      <div className="flex justify-center">
        <div className="pointer-events-auto w-full max-w-4xl mx-auto px-4 sm:px-6 py-4 transition-all duration-300 rounded-full bg-gray-50/50 border border-gray-300/40 backdrop-blur-md shadow-sm">
          <nav className="flex items-center justify-between">
            <Link href="/" className="flex items-center space-x-2">
              <div className="bg-gradient-to-r from-blue-600 to-cyan-600 p-2 rounded-lg">
                <LinkIcon className="w-5 h-5 sm:w-6 sm:h-6 text-white" />
              </div>
              <span className="text-lg sm:text-xl font-bold text-black">ShortURL</span>
            </Link>
          </nav>
        </div>
      </div>
    </header>
  )
}

export default Header