#!/bin/bash
output=$1

check_result() {
  local success_msg=$1
  local failure_msg=$1

  if [ $? -eq 0 ]; then
    echo $success_msg
  else
    echo $failure_msg
    exit 1
  fi
}

# container changing 
./iproc.exe -i images/almoço.png -o $output/container-change.jpg
check_result "container changing success" "error on container changing"

# blur
./iproc.exe -i images/angra.jpg -o $output/blur.jpg -b 11 -s 5
check_result "blur success" "error on blurring image"

# brightness
./iproc.exe -i images/AngraAngelsCry.jpg -o $output/brightness.jpg -l 20
check_result "brightness success" "error on blurring image"

# crop
./iproc.exe -i images/art\ vs\ science.jpg -o $output/crop.jpg -c 100,100
check_result "crop success" "error on cropping image"

# flix x
./iproc.exe -i images/AwakenMyLove.png -o $output/flip-x.jpg -fx
check_result "flip x success" "error on flip x"

# flix y
./iproc.exe -i images/Capa_de_Músicas_para_Churrasco,_Vol._1.jpg -o $output/flip-y.jpg -fy
check_result "flip y success" "error on flip y"

# grayscale
./iproc.exe -i images/imprevisto.jpg -o $output/grayscale.jpg -gs
check_result "grayscale success" "error on grayscale"

# resize factor (nearest neighbor)
./iproc.exe -i images/SamSpratt_KidCudi_ManOnTheMoon3_AlbumCover_Web.jpg -o $output/resize-nn-f.jpg -nn -f 1.5
check_result "resize factor (nearest neighbor) success" "error on resize factor (nearest neighbor)"

# resize with target (nearest neighbor)
./iproc.exe -i images/SamSpratt_KidCudi_ManOnTheMoon3_AlbumCover_Web.jpg -o $output/resize-nn-w-h.jpg -nn -width 2000 -height 1000
check_result "resize with target (nearest neighbor) success" "error on resize factor (nearest neighbor)"

# overlay
./iproc.exe -i images/almoço.png -ov images/sobrevivendo_no_inferno.jpg -o $output/overlay.jpg
check_result "overlay success" "error on overlay"


echo "sucessful test running - please check the results on output folder"
exit 0