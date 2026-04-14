import type { Metadata } from 'next'
import '../styles/globals.css'
import { Toaster } from '@/components/ui/toaster'

export const metadata: Metadata = {
  title: 'URL Shortener',
  description: 'A simple and fast URL shortener service',
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className="bg-gray-50 text-gray-900">
        {children}
        <Toaster />
      </body>
    </html>
  )
}