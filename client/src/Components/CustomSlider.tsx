import { useState, useEffect } from "react"

const images = ["/animate.png", "/animate2.png", "/animate3.png", "/animate4.png", "/animate5.png"]

function CustomSlider() {
  const [index, setIndex] = useState(0)
   
  useEffect(() => {
    const interval = setInterval(() => {
      setIndex((prevIndex) => (prevIndex === images.length -1 ? -1  : prevIndex +1))
    }, 2000)

    return () => clearInterval(interval)
  }, [])

  return (
    <div className="relative w-72 h-72 overflow-hidden">
      <div
        className="flex transition-transform duration-700 ease-in-out"
        style={{ transform: `translateX(-${index * 100}%)` }}
      >
        {images.map((src, i) => (
          <img key={i} src={src} className="w-72 flex-shrink-0" alt="Slide" />
        ))}
      </div>
    </div>
  )
}

export default CustomSlider


